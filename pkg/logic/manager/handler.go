package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"sync"
)

type clientHandler func(req *proto.Request, cid int32) (res []byte, err error)

type handler struct {
	sync.Once
	id             int32
	router         string
	managerHandler clientHandler
	removeChan     chan<- int32
	outputChan     chan<- []byte
	pingFunc       func()
	closeFunc      func()
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewHandler(ctx context.Context, router string, mh clientHandler, removeChan chan<- int32) *handler {
	sub, cancel := context.WithCancel(ctx)
	c := &handler{
		id:             0,
		router:         router,
		managerHandler: mh,
		removeChan:     removeChan,
		ctx:            sub,
		cancel:         cancel,
	}
	return c
}

func (c *handler) RegisterId(id int32) {
	c.id = id
}

func (c *handler) RegisterRemoveChan(ch chan<- int32) {
	c.removeChan = ch
}

func (c *handler) RegisterConnWriteChan(ch chan<- []byte) {
	c.outputChan = ch
}

func (c *handler) RegisterConnClose(do func()) {
	c.closeFunc = do
}

func (c *handler) RegisterConnPing(do func()) {
	c.pingFunc = do
}

func (c *handler) Handler(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIPing:
		c.pingFunc()
	default:
		return c.managerHandler(req, c.id)
	}
	return res, nil
}

func (c *handler) Close() {
	c.Once.Do(func() {
		c.removeChan <- c.id
		c.cancel()
	})
}
