package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"k8s.io/klog/v2"
	"sync"
)

func newGatewayHub(ctx context.Context) *gatewayHub {
	return &gatewayHub{
		connections: websocket.NewConnections(ctx),
		items:       make(map[int32]*proto.Gateway, 0),
		clearChan:   make(chan int32, 1024),
		ctx:         ctx,
	}
}

type gatewayHub struct {
	sync.RWMutex
	connections *websocket.Connections
	items       map[int32]*proto.Gateway
	clearChan   chan int32
	ctx         context.Context
}

func (m *Manager) loopClearGateway() {
	for {
		select {
		case <-m.gateways.ctx.Done():
			return
		case id, isClose := <-m.gateways.clearChan:
			if !isClose {
				return
			}
			m.gateways.Lock()
			delete(m.gateways.items, id)
			m.gateways.Unlock()
			// todo push to all
		}
	}
}

func (m *Manager) handlerGateway(in []byte, handler handler.Interface) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIPing:
		handler.Ping()
	case proto.ServiceAPIGatewayRegister:
		ga := &proto.Gateway{}
		if err = ga.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.gateways.Lock()
		m.gateways.items[handler.Id()] = ga
		m.gateways.Unlock()
		klog.Infof("handlerGateway items:%v", m.gateways.items)
		// todo push to all
		if err = m.updateWorkerList(); err != nil {
			klog.V(2).Info(err)
			return res, err
		}
	}
	return res, nil
}

func (m *Manager) updateGatewayList() error {
	m.gateways.RLock()
	m.gateways.RUnlock()
	return nil
}
