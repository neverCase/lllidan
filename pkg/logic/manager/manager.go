package manager

import (
	"context"
	cmap "github.com/nevercase/concurrent-map"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
)

type Manager struct {
	gateways   *gatewayHub
	workers    *websocket.Connections
	dashboards *websocket.Connections
	handlers   cmap.ConcurrentMap
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		gateways: &gatewayHub{
			connections: websocket.NewConnections(ctx),
			items:       make(map[string]*proto.Gateway, 0),
		},
		workers:    websocket.NewConnections(ctx),
		dashboards: websocket.NewConnections(ctx),
		handlers:   cmap.New(),
	}
	return m
}

func (m *Manager) handlerWorker(req *proto.Request) (res []byte, err error) {

	return res, nil
}

func (m *Manager) handlerDashboard(req *proto.Request) (res []byte, err error) {
	return res, nil
}

func (m *Manager) Handler() handler.Interface {
	return &client{}
}
