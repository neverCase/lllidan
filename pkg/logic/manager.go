package logic

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"k8s.io/klog/v2"
	"sync"
)

type Manager struct {
	gateways   *websocket.Connections
	workers    *websocket.Connections
	dashboards *websocket.Connections
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		gateways:   websocket.NewConnections(ctx),
		workers:    websocket.NewConnections(ctx),
		dashboards: websocket.NewConnections(ctx),
	}
	return m
}

func (m *Manager) handlerGateway(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return res, nil
}

func (m *Manager) handlerWorker(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return res, nil
}

func (m *Manager) handlerDashboard(in []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return res, nil
}

func (m *Manager) Handler() handler.Handler {
	return &client{}
}

type client struct {
	sync.Once
	router     string
	clearChan  chan<- int32
	outputChan chan<- []byte
	pingFunc   func()
	closeFunc  func()
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewClient(ctx context.Context, router string) *client {
	sub, cancel := context.WithCancel(ctx)
	c := &client{
		router: router,
		ctx:    sub,
		cancel: cancel,
	}
	return c
}

func (c *client) Close() {
	c.Once.Do(func() {

	})
}

func (c *client) Handler(data []byte) (res []byte, err error) {

	return res, nil
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
