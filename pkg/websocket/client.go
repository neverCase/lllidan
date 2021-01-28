package websocket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
	"net/url"
	"sync"
	"time"
)

const (
	ConnectionTimeoutInSec      = 10
	WriteToChannelTimeoutInMS   = 1000
	RetryConnectionDurationInMS = 3000
)

func NewOption(ctx context.Context, hostname string, address string, path string) *Option {
	sub, cancel := context.WithCancel(ctx)
	return &Option{
		Hostname:          hostname,
		Address:           address,
		Path:              path,
		Status:            OptionInActive,
		registerFunctions: make([]OptionRegisterFunction, 0),
		readHandlerChan:   make(chan []byte, 4096),
		writeHandlerChan:  make(chan []byte, 4096),
		ctx:               sub,
		cancelFunc:        cancel,
		writeTimeout:      WriteToChannelTimeoutInMS,
		retryDuration:     RetryConnectionDurationInMS,
	}
}

type OptionStatus string

const (
	OptionActive   OptionStatus = "active"
	OptionInActive OptionStatus = "inactive"
	OptionClosed   OptionStatus = "closed"
)

type OptionRegisterFunction func() error

type Option struct {
	rw                sync.RWMutex
	once              sync.Once
	Hostname          string
	Address           string
	Path              string
	Status            OptionStatus
	registerFunctions []OptionRegisterFunction
	readHandlerChan   chan []byte
	writeHandlerChan  chan []byte
	ctx               context.Context
	cancelFunc        context.CancelFunc
	retryDuration     int64
	writeTimeout      int64
}

func (o *Option) RegisterFunc(do ...OptionRegisterFunction) {
	for _, v := range do {
		o.registerFunctions = append(o.registerFunctions, v)
	}
}

func (o *Option) Prepare() error {
	for _, v := range o.registerFunctions {
		if err := v(); err != nil {
			klog.V(2).Infof("Prepare do func err:%v", err)
			return err
		}
	}
	return nil
}

func (o *Option) Read() <-chan []byte {
	return o.readHandlerChan
}

func (o *Option) Send(in []byte) error {
	o.rw.RLock()
	defer o.rw.RUnlock()
	switch o.Status {
	case OptionActive:
		select {
		case <-time.After(time.Millisecond * time.Duration(o.writeTimeout)):
			return fmt.Errorf("option send message failed due to write to channel timeout:%v ms", o.writeTimeout)
		case o.writeHandlerChan <- in:
		}
	case OptionInActive, OptionClosed:
		return fmt.Errorf("option skip sending message due to the status was:%v", o.Status)
	}
	return nil
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
	tick := time.NewTicker(time.Second * time.Duration(ConnectionTimeoutInSec/2))
	defer tick.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-tick.C:
			c.opt.writeHandlerChan <- in
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
		c.opt.readHandlerChan <- message
	}
}

func (c *Client) writePump() {
	defer c.Close()
	for {
		select {
		case msg, isClose := <-c.opt.writeHandlerChan:
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
