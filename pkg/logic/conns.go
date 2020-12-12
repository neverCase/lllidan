package logic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	cmap "github.com/nevercase/concurrent-map"
	"k8s.io/klog/v2"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	WebsocketConnectionTimeout = 10
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewConnections(ctx context.Context) *connections {
	cs := &connections{
		autoIncr:         0,
		dashboardClients: cmap.New(),
		gatewayClients:   cmap.New(),
		removedChan:      make(chan *removeConn, 4096),
		ctx:              ctx,
	}
	go cs.remove()
	return cs
}

type connections struct {
	autoIncr         int32
	dashboardClients cmap.ConcurrentMap
	gatewayClients   cmap.ConcurrentMap
	removedChan      chan *removeConn
	ctx              context.Context
}

type connType int

const (
	connTypeDashboard connType = iota
	connTypeGateway   connType = 1
)

type removeConn struct {
	id       int32
	connType connType
}

func (cs *connections) remove() {
	for {
		select {
		case <-cs.ctx.Done():
			return
		case o, isClose := <-cs.removedChan:
			if !isClose {
				return
			}
			shardKey := shardKey(o.id)
			var (
				obj interface{}
				ok  bool
			)
			switch o.connType {
			case connTypeDashboard:
				obj, ok = cs.dashboardClients.Get(shardKey)
			case connTypeGateway:
				obj, ok = cs.gatewayClients.Get(shardKey)
			}
			if !ok {
				continue
			}
			t := obj.(cmap.ConcurrentMap)
			t.Remove(shardKey)
		}
	}
}

func shardKey(id int32) string {
	return fmt.Sprintf("%d", id)
}

func (cs *connections) handler(w http.ResponseWriter, r *http.Request, connType connType, handler connHandler) {
	client, err := cs.newConn(w, r, connType, handler)
	if err != nil {
		klog.V(2).Info(err)
		return
	}
	shardKey := shardKey(client.id)
	var obj interface{}
	switch connType {
	case connTypeDashboard:
		cs.dashboardClients.SetIfAbsent(shardKey, cmap.New())
		obj, _ = cs.dashboardClients.Get(shardKey)
	case connTypeGateway:
		cs.gatewayClients.SetIfAbsent(shardKey, cmap.New())
		obj, _ = cs.gatewayClients.Get(shardKey)
	}
	t := obj.(cmap.ConcurrentMap)
	t.Set(shardKey, client)
}

func (cs *connections) newConn(w http.ResponseWriter, r *http.Request, connType connType, handler connHandler) (*conn, error) {
	client, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	ctx, cancel := context.WithCancel(cs.ctx)
	c := &conn{
		id:                    atomic.AddInt32(&cs.autoIncr, 1),
		conn:                  client,
		connType:              connType,
		handler:               handler,
		writeChan:             make(chan []byte, 4096),
		lastPingTime:          time.Now(),
		keepAliveTimeoutInSec: WebsocketConnectionTimeout,
		closeOnce:             sync.Once{},
		removedChan:           cs.removedChan,
		ctx:                   ctx,
		cancel:                cancel,
	}
	go c.keepAlive()
	go c.readPump()
	go c.writePump()
	return c, nil
}

type connHandler func(data []byte, do func(), outputChan chan<- []byte) (res []byte, err error)

type conn struct {
	id                    int32
	conn                  *websocket.Conn
	connType              connType
	handler               connHandler
	writeChan             chan []byte
	lastPingTime          time.Time
	keepAliveTimeoutInSec int64
	closeOnce             sync.Once
	removedChan           chan<- *removeConn
	ctx                   context.Context
	cancel                context.CancelFunc
}

func (c *conn) keepAlive() {
	defer c.close()
	tick := time.NewTicker(time.Second * time.Duration(c.keepAliveTimeoutInSec+1))
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if time.Now().Sub(c.lastPingTime) > time.Second*time.Duration(c.keepAliveTimeoutInSec) {
				klog.Info("keepAlive timeout")
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *conn) ping() {
	c.lastPingTime = time.Now()
}

func (c *conn) close() {
	c.closeOnce.Do(func() {
		c.cancel()
		c.removedChan <- &removeConn{id: c.id, connType: c.connType}
		if err := c.conn.Close(); err != nil {
			klog.V(2).Info(err)
		}
	})
}

func (c *conn) readPump() {
	defer c.close()
	for {
		messageType, data, err := c.conn.ReadMessage()
		klog.V(3).Infof("messageType: %d message-string: %s\n", messageType, string(data))
		if err != nil {
			klog.V(2).Info(err)
			return
		}
		res, err := c.handler(data, c.ping, c.writeChan)
		if err != nil {
			klog.V(2).Info(err)
			//return
		}
		if len(res) == 0 {
			klog.V(3).Info("ws conn handler res len 0")
			continue
		}
		c.writeChan <- res
	}
}

func (c *conn) writePump() {
	defer c.close()
	for {
		select {
		case msg, isClose := <-c.writeChan:
			if !isClose {
				return
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				klog.V(2).Info(err)
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}