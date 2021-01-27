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

func NewOption(ctx context.Context, hostname string, address string, path string) *Option {
	sub, cancel := context.WithCancel(ctx)
	return &Option{
		Hostname:         hostname,
		Address:          address,
		Path:             path,
		ReadHandlerChan:  make(chan []byte, 4096),
		WriteHandlerChan: make(chan []byte, 4096),
		ctx:              sub,
		cancelFunc:       cancel,
	}
}

type Option struct {
	once             sync.Once
	Hostname         string
	Address          string
	Path             string
	ReadHandlerChan  chan []byte
	WriteHandlerChan chan []byte
	ctx              context.Context
	cancelFunc       context.CancelFunc
}

func (o *Option) Done() <-chan struct{} {
	return o.ctx.Done()
}

func (o *Option) Cancel() {
	o.once.Do(func() {
		o.cancelFunc()
	})
}

func NewClient(opt *Option) (*Client, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   opt.Address,
		Path:   opt.Path,
	}
	a, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	c := &Client{
		opt:  opt,
		conn: a,
	}
	go c.readPump()
	go c.writePump()
	return c, nil
}

type Client struct {
	sync.Once
	opt    *Option
	conn   *websocket.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Client) Option() *Option {
	return c.opt
}

func (c *Client) Close() {
	c.Once.Do(func() {
		c.opt.cancelFunc()
	})
}

func (c *Client) Ping(in []byte) {
	defer c.Close()
	tick := time.NewTicker(time.Second * time.Duration(logic.WebsocketConnectionTimeout/2))
	defer tick.Stop()
	for {
		select {
		case <-c.opt.ctx.Done():
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
		case <-c.opt.ctx.Done():
			return
		}
	}
}
