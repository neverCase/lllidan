package manager

import (
	"context"
	"fmt"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/env"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"net/http"
	"net/url"
)

type Manager struct {
	conf     *config.Config
	hostname string
	logic    *logic
	gateways *gatewayHub
}

func NewManager(ctx context.Context, conf *config.Config) (*Manager, error) {
	hostname, err := env.GetHostName()
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", conf.Logic.KubernetesService.Name, conf.Logic.KubernetesService.Port),
		Path:   proto.RouterWorker,
	}
	m := &Manager{
		conf:     conf,
		hostname: hostname,
		gateways: newGatewayHub(ctx),
		logic:    InitLogic(ctx, u, hostname),
	}
	// start logic
	m.logic.option.RegisterFunc(m.registerWorker)
	go websocket.NewClientWithReconnect(m.logic.option)
	return m, nil
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request, router string) {
	switch router {
	case proto.RouterGateway:
		m.gateways.connections.Handler(w, r, handler.NewHandler(context.Background(), router, m.handlerGateway, m.gateways.clearChan))
	}
	return
}
