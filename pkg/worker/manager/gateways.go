package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"k8s.io/klog/v2"
	"sync"
	"sync/atomic"
)

func newGatewayHub(ctx context.Context) *gatewayHub {
	gh := &gatewayHub{
		connections: websocket.NewConnections(ctx),
		items:       make(map[int32]*gateway, 0),
		clearChan:   make(chan int32, 1024),
		messageHandler: &MessageOption{
			ReadChan:         make(chan []byte, 4096),
			WriteChan:        make(chan []byte, 4096),
			writeTimeoutInMs: 1000,
		},
		ctx: ctx,
	}
	go gh.pushMessage()
	return gh
}

type gatewayHub struct {
	sync.RWMutex
	connections    *websocket.Connections
	items          map[int32]*gateway
	clearChan      chan int32
	messageHandler *MessageOption
	ctx            context.Context
}

type gateway struct {
	gateway    *proto.Gateway
	sendNumber int32
}

func (gh *gatewayHub) GatewayMessageHandler() handler.MessageHandler {
	return gh.messageHandler
}

func (gh *gatewayHub) newGateway(g *proto.Gateway) *gateway {
	return &gateway{
		gateway:    g,
		sendNumber: 0,
	}
}

func (gh *gatewayHub) resetLoadBalanceNumber() {
	gh.Lock()
	defer gh.Unlock()
	for _, v := range gh.items {
		v.sendNumber = 0
	}
}

func (gh *gatewayHub) loadBalancePush(in []byte) bool {
	gh.Lock()
	defer gh.Unlock()
	var id int32
	for k, v := range gh.items {
		if id == 0 || v.sendNumber < gh.items[id].sendNumber {
			id = k
		}
	}
	if id == 0 {
		return false
	}
	atomic.AddInt32(&gh.items[id].sendNumber, 1)
	return gh.connections.SendToOne(in, id)
}

func (gh *gatewayHub) pushMessage() {
	for {
		select {
		case msg, isClose := <-gh.messageHandler.WriteChan:
			if !isClose {
				return
			}
			switch gh.loadBalancePush(msg) {
			case true:
				klog.Info("gatewayHub pushMessage success")
			case false:
				klog.Info("gatewayHub pushMessage failed")
			}
		}
	}
}

func (m *Manager) loopClearGateway() {
	for {
		select {
		case <-m.gateways.ctx.Done():
			return
		case id, isClose := <-m.gateways.clearChan:
			if !isClose {
				return
			}
			m.gateways.Lock()
			delete(m.gateways.items, id)
			m.gateways.Unlock()
		}
	}
}

func (m *Manager) handlerGateway(req *proto.Request, id int32) (res []byte, err error) {
	switch req.ServiceAPI {
	case proto.ServiceAPIGatewayConnect:
		ga := &proto.Gateway{}
		if err = ga.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.gateways.Lock()
		m.gateways.items[id] = m.gateways.newGateway(ga)
		m.gateways.Unlock()
		klog.Infof("handlerGateway items:%v", m.gateways.items)
		// reset
		m.gateways.resetLoadBalanceNumber()
	default:
		var in []byte
		if in, err = req.Marshal(); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		if err = m.gateways.messageHandler.writeToReadChan(in); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
	}
	return res, nil
}

func (m *Manager) GatewayMessageHandler() handler.MessageHandler {
	return m.gateways.messageHandler
}
