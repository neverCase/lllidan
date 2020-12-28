package gateway

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	cmap "github.com/nevercase/concurrent-map"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
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
		autoIncr:    0,
		members:     cmap.New(),
		removedChan: make(chan int32, 4096),
		ctx:         ctx,
	}
	go cs.remove()
	return cs
}

type connections struct {
	autoIncr    int32
	members     cmap.ConcurrentMap
	removedChan chan int32
	ctx         context.Context
}

func (cs *connections) remove() {
	for {
		select {
		case <-cs.ctx.Done():
			return
		case id, isClose := <-cs.removedChan:
			if !isClose {
				return
			}
			shardKey := shardKey(id)
			obj, ok := cs.members.Get(shardKey)
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

func (cs *connections) handler(w http.ResponseWriter, r *http.Request, url url.URL) {
	opt := &ProxyOption{
		Url:          url,
		IsMaintained: true,
		Id:           atomic.AddInt32(&cs.autoIncr, 1),
	}
	client, err := cs.newConn(w, r, opt)
	if err != nil {
		klog.V(2).Info(err)
		return
	}
	shardKey := shardKey(client.handler.Id())
	cs.members.SetIfAbsent(shardKey, cmap.New())
	obj, _ := cs.members.Get(shardKey)
	t := obj.(cmap.ConcurrentMap)
	t.Set(shardKey, client)
}

func (cs *connections) newConn(w http.ResponseWriter, r *http.Request, opt *ProxyOption) (*conn, error) {
	client, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	p, err := NewProxy(context.Background(), opt)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(cs.ctx)
	c := &conn{
		conn:                  client,
		handler:               p,
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

type conn struct {
	conn                  *websocket.Conn
	handler               Proxy
	lastPingTime          time.Time
	keepAliveTimeoutInSec int64
	closeOnce             sync.Once
	removedChan           chan<- int32
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
		go c.handler.Close()
		if err := c.conn.Close(); err != nil {
			klog.V(2).Info(err)
		}
		c.cancel()
		c.removedChan <- c.handler.Id()
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
		if err = c.handler.HandlerWrite(data); err != nil {
			klog.V(2).Info(err)
		}
	}
}

func (c *conn) writePump() {
	defer c.close()
	for {
		select {
		case msg, isClose := <-c.handler.HandlerRead():
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
