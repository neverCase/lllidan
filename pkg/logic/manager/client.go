package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"sync"
)

type clientHandler func(req *proto.Request) (res []byte, err error)

type client struct {
	sync.Once
	router         string
	managerHandler clientHandler
	clearChan      chan<- int32
	outputChan     chan<- []byte
	pingFunc       func()
	closeFunc      func()
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewClient(ctx context.Context, router string, mh clientHandler) *client {
	sub, cancel := context.WithCancel(ctx)
	c := &client{
		router:         router,
		managerHandler: mh,
		ctx:            sub,
		cancel:         cancel,
	}
	return c
}

func (c *client) RegisterConnWriteChan(ch chan<- []byte) {
	c.outputChan = ch
}

func (c *client) RegisterConnClose(do func()) {
	c.closeFunc = do
}

func (c *client) RegisterConnPing(do func()) {
	c.pingFunc = do
}

func (c *client) Handler(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIPing:
		c.pingFunc()
	default:
		return c.managerHandler(req)
	}
	return res, nil
}

func (c *client) Close() {
	c.Once.Do(func() {
	})
}

