package manager

import (
	"context"
	cmap "github.com/nevercase/concurrent-map"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"net/http"
)

type Manager struct {
	gateways   *gatewayHub
	workers    *workerHub
	dashboards *websocket.Connections
	handlers   cmap.ConcurrentMap
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		gateways:   newGatewayHub(ctx),
		workers:    newWorkerHub(ctx),
		dashboards: websocket.NewConnections(ctx),
		handlers:   cmap.New(),
	}
	go m.loopClearGateway()
	return m
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request, router string) {
	switch router {
	case proto.RouterGateway:
		m.gateways.connections.Handler(w, r, handler.NewHandler(context.Background(), router, m.handlerGateway, m.gateways.clearChan))
	case proto.RouterWorker:
		m.workers.connections.Handler(w, r, handler.NewHandler(context.Background(), router, m.handlerWorker, m.workers.clearChan))
	case proto.RouterDashboard:
	}
	return
}
