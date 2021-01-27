package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
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

func (m *Manager) handlerGateway(req *proto.Request, id int32) (res []byte, err error) {
	switch req.ServiceAPI {
	case proto.ServiceAPIGatewayRegister:
		ga := &proto.Gateway{}
		if err = ga.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.gateways.Lock()
		m.gateways.items[id] = ga
		m.gateways.Unlock()
		// todo push to all
	}
	return res, nil
}

func (m *Manager) updateGatewayList() error {
	m.gateways.RLock()
	m.gateways.RUnlock()
	return nil
}