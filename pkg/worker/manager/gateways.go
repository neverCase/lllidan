package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"sync"
)

func newGatewayHub(ctx context.Context) *gatewayHub {
	return &gatewayHub{
		connections: websocket.NewConnections(ctx),
		items:       make(map[int32]*gateway, 0),
		clearChan:   make(chan int32, 1024),
		ctx:         ctx,
	}
}

type gatewayHub struct {
	sync.RWMutex
	connections *websocket.Connections
	items       map[int32]*gateway
	clearChan   chan int32
	ctx         context.Context
}

type gateway struct {
	gateway    *proto.Gateway
	sendNumber int32
}
