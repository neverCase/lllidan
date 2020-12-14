package logic

import (
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"sync"
)

type manager struct {
	mu            sync.RWMutex
	gatewayIncrId int32
	gateways      map[string]proto.Gateway
}

func newManager() *manager {
	return &manager{
		gatewayIncrId: 0,
		gateways:      make(map[string]proto.Gateway, 0),
	}
}

func (m *manager) gatewayRegister(gateway *proto.Gateway) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.gateways[gateway.Hostname]; ok {

	} else {
		m.gateways[gateway.Hostname] = *gateway
	}
}

func (m *manager) listGateway() {

}

func (s *Server) handlerDashboard(data []byte, ping func(), outputChan chan<- []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(data); err != nil {
		klog.V(2).Info(err)
	}
	switch req.ServiceAPI {
	case proto.APIPing:
		ping()
	}
	return res, nil
}

func (s *Server) handlerGateway(data []byte, ping func(), outputChan chan<- []byte) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(data); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.APIPing:
		ping()
	case proto.APIGatewayRegister:
		obj := &proto.Gateway{}
		if err = obj.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		s.manager.gatewayRegister(obj)
		// todo sync to all the gateway clients

		// todo sync to all the dashboards
	}
	return res, nil
}
