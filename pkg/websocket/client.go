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
		Status:           OptionInActive,
		ctx:              sub,
		cancelFunc:       cancel,
	}
}

type OptionStatus string

const (
	OptionActive   OptionStatus = "active"
	OptionInActive OptionStatus = "inactive"
	OptionClosed   OptionStatus = "closed"
)

type Option struct {
	rw               sync.RWMutex
	once             sync.Once
	Hostname         string
	Address          string
	Path             string
	ReadHandlerChan  chan []byte
	WriteHandlerChan chan []byte
	Status           OptionStatus
	ctx              context.Context
	cancelFunc       context.CancelFunc
}

func (o *Option) Send(in []byte) {
	o.rw.RLock()
	defer o.rw.RUnlock()
	switch o.Status {
	case OptionActive:
		o.WriteHandlerChan <- in
	case OptionInActive, OptionClosed:
		klog.Infof("option skip sending message due to the status was:v", o.Status)
	}
}

func (o *Option) ChangeStatus(s OptionStatus) {
	o.rw.Lock()
	defer o.rw.Unlock()
	switch o.Status {
	case OptionClosed:
	case OptionActive, OptionInActive:
		o.Status = s
	}
}

func (o *Option) Done() <-chan struct{} {
	return o.ctx.Done()
}

func (o *Option) Cancel() {
	o.once.Do(func() {
		o.ChangeStatus(OptionClosed)
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
	sub, cancel := context.WithCancel(opt.ctx)
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
