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
	workers  *workerHub
	apis     *apiHub
	logic    *logic
}

func NewManger(ctx context.Context, conf *config.Config) (*Manager, error) {
	hostname, err := env.GetHostName()
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", conf.Logic.KubernetesService.Name, conf.Logic.KubernetesService.Port),
		Path:   proto.RouterGateway,
	}
	req, err := GetConnectToWorkerRequest(hostname, int32(conf.Gateway.ListenPort))
	if err != nil {
		return nil, err
	}
	m := &Manager{
		conf:     conf,
		hostname: hostname,
		workers:  NewWorkerHub(ctx, hostname, req),
		apis:     NewApiHub(ctx),
		logic:    InitLogic(ctx, u, hostname),
	}
	// start logic
	m.logic.option.RegisterFunc(m.registerGateway)
	go websocket.NewClientWithReconnect(m.logic.option)
	// handler logic message
	go m.loopLogicMessage(ctx)
	return m, nil
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request, router string) {
	switch router {
	case proto.RouterGatewayApi:
		m.apis.connections.Handler(w, r, handler.NewHandler(context.Background(), router, m.handlerApi, m.apis.clearChan))
	}
	return
}