package gateway

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/nevercase/lllidan/pkg/logic"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"net/url"
	"sync"
	"time"
)

func NewClientWithRecover(addr string, stop <-chan struct{}, handler chan<- []byte) (*Client, error) {
	for {
		client, err := NewClient(addr, handler)
		if err != nil {
			klog.V(2).Info(err)
			//continue
			return nil, err
		}
		select {
		case <-client.Done():
		case <-stop:
			return nil, nil
		}
	}
}

func NewClient(addr string, handler chan<- []byte) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: proto.RouterGateway}
	a, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{
		conn:        a,
		writeChan:   make(chan []byte, 1024),
		handlerChan: handler,
		ctx:         ctx,
		cancel:      cancel,
	}
	if err = c.registerGateway(); err != nil {
		return nil, err
	}
	go c.readPump()
	go c.writePump()
	go c.ping()
	return c, nil
}

type Client struct {
	conn        *websocket.Conn
	writeChan   chan []byte
	handlerChan chan<- []byte
	once        sync.Once
	ctx         context.Context
	cancel      context.CancelFunc
}

func (c *Client) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Client) close() {
	c.once.Do(func() {
		c.cancel()
	})
}

func (c *Client) ping() {
	defer c.close()
	tick := time.NewTicker(time.Second * time.Duration(logic.WebsocketConnectionTimeout/2))
	defer tick.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-tick.C:
			req := &proto.Request{
				ServiceAPI: proto.APIPing,
				Data:       make([][]byte, 0),
			}
			data, err := req.Marshal()
			if err != nil {
				klog.V(2).Info(err)
				return
			}
			c.writeChan <- data
		}
	}
}

func (c *Client) readPump() {
	defer c.close()
	for {
		messageType, message, err := c.conn.ReadMessage()
		klog.V(3).Infof("messageType: %d message: %s err:%v\n", messageType, string(message), err)
		if err != nil {
			klog.V(2).Info(err)
			return
		}
		c.handlerChan <- message
	}
}

func (c *Client) writePump() {
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

func (c *Client) registerGateway() (err error) {
	gw := &proto.Gateway{
		Hostname: "",
		Ip:       "",
		Port:     0,
	}
	var data []byte
	if data, err = gw.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	}
	req := &proto.Request{
		ServiceAPI: proto.APIGatewayRegister,
		Data:       make([][]byte, 0),
	}
	req.Data = append(req.Data, data)
	if data, err = req.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	}
	c.writeChan <- data
	return nil
}
