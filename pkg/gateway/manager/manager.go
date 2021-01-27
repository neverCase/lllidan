package manager

import (
	"context"
	cmap "github.com/nevercase/concurrent-map"
	"github.com/nevercase/lllidan/pkg/env"
)

type Manager struct {
	workers  *workerHub
	handlers cmap.ConcurrentMap
}

func NewManger(ctx context.Context) (*Manager, error) {
	hostname, err := env.GetHostName()
	if err != nil {
		return nil, err
	}
	return &Manager{
		workers: NewWorkerHub(ctx, hostname),
	}, nil
}