package manager

import (
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"k8s.io/klog/v2"
	"sync"
)

type gatewayHub struct {
	sync.RWMutex
	connections *websocket.Connections
	items       map[string]*proto.Gateway
}

func (m *Manager) handlerGateway(req *proto.Request) (res []byte, err error) {
	switch req.ServiceAPI {
	case proto.ServiceAPIGatewayRegister:
		ga := &proto.Gateway{}
		if err = ga.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.gateways.Lock()
		m.gateways.items[ga.Hostname] = ga
		m.gateways.Unlock()
		// todo push to all dashboard
	}
	return res, nil
}
