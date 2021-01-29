package handler

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

func (h *handler) RegisterId(id int32) {
	h.id = id
}

func (h *handler) RegisterRemoveChan(ch chan<- int32) {
	h.removeChan = ch
}

func (h *handler) RegisterConnWriteChan(ch chan<- []byte) {
	h.outputChan = ch
}

func (h *handler) RegisterConnClose(do func()) {
	h.closeFunc = do
}

func (h *handler) RegisterConnPing(do func()) {
	h.pingFunc = do
}

func (h *handler) Handler(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIPing:
		h.pingFunc()
	default:
		return h.managerHandler(req, h.id)
	}
	return res, nil
}

func (h *handler) Close() {
	h.Once.Do(func() {
		h.removeChan <- h.id
		h.cancel()
	})
}

