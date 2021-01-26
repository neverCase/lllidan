package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"sync"
)

func newWorkerHub(ctx context.Context) *workerHub {
	return &workerHub{
		connections: websocket.NewConnections(ctx),
		items:       make(map[string]*proto.Worker, 0),
	}
}

type workerHub struct {
	sync.RWMutex
	connections *websocket.Connections
	items       map[string]*proto.Worker
}

func (m *Manager) handlerWorker(req *proto.Request, id int32) (res []byte, err error) {
	return res, nil
}
