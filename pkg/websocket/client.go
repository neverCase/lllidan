package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/nevercase/lllidan/pkg/logic"
	"k8s.io/klog/v2"
	"net/url"
	"sync"
	"time"
)

type Option struct {
	Hostname         string
	Address          string
	Path             string
	ReadHandlerChan  chan []byte
	WriteHandlerChan chan []byte
}

func NewClient(ctx context.Context, opt *Option) (*Client, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   opt.Address,
		Path:   opt.Path,
	}
	a, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	sub, cancel := context.WithCancel(ctx)
	c := &Client{
		opt:    opt,
		conn:   a,
		ctx:    sub,
		cancel: cancel,
	}
	go c.readPump()
	go c.writePump()
	return c, nil
}

type Client struct {
	sync.Once
	opt         *Option
	conn        *websocket.Conn
	ctx         context.Context
	cancel      context.CancelFunc
}

func (c *Client) Close() {
	c.Once.Do(func() {
		c.cancel()
	})
}

func (c *Client) Ping(in []byte) {
	defer c.Close()
	tick := time.NewTicker(time.Second * time.Duration(logic.WebsocketConnectionTimeout/2))
	defer tick.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-tick.C:
			c.opt.WriteHandlerChan <- in
		}
	}
}

func (c *Client) readPump() {
	defer c.Close()
	for {
		_, message, err := c.conn.ReadMessage()
		//klog.V(3).Infof("messageType: %d message: %s err:%v\n", messageType, string(message), err)
		if err != nil {
			klog.V(2).Info(err)
			return
		}
		c.opt.ReadHandlerChan <- message
	}
}

func (c *Client) writePump() {
	defer c.Close()
	for {
		select {
		case msg, isClose := <-c.opt.WriteHandlerChan:
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
