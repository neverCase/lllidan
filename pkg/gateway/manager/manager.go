package manager

import (
	"context"
	"fmt"
	cmap "github.com/nevercase/concurrent-map"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/env"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"net/url"
)

type Manager struct {
	conf     *config.Config
	hostname string
	workers  *workerHub
	logic    *logic
	handlers cmap.ConcurrentMap
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
	m := &Manager{
		conf:     conf,
		hostname: hostname,
		workers:  NewWorkerHub(ctx, hostname),
		logic:    InitLogic(ctx, u, hostname),
	}
	// start logic
	m.logic.option.RegisterFunc(m.registerGateway)
	go websocket.NewClientWithReconnect(m.logic.option)
	// handler logic message
	go m.loopLogicMessage(ctx)
	return m, nil
}
