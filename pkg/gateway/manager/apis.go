package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"sync"
)

type apiHub struct {
	sync.RWMutex
	connections *websocket.Connections
	clearChan   chan int32
	ctx         context.Context
}

func NewApiHub(ctx context.Context) *apiHub {
	a := &apiHub{
		connections: websocket.NewConnections(ctx),
		clearChan:   make(chan int32, 1),
		ctx:         ctx,
	}
	go a.loopClearApi()
	return a
}

func (a *apiHub) loopClearApi() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case _, isClose := <-a.clearChan:
			if !isClose {
				return
			}
		}
	}
}

func (m *Manager) handlerApi(req *proto.Request, id int32) (res []byte, err error) {
	var data []byte
	if data, err = req.Marshal(); err != nil {
		return nil, err
	}
	m.workers.readChan <- data
	return res, nil
}
