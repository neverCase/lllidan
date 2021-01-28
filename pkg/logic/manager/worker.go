package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"k8s.io/klog/v2"
	"sync"
)

func newWorkerHub(ctx context.Context) *workerHub {
	return &workerHub{
		connections: websocket.NewConnections(ctx),
		items:       make(map[int32]*proto.Worker, 0),
		clearChan:   make(chan int32, 1024),
		ctx:         ctx,
	}
}

type workerHub struct {
	sync.RWMutex
	connections *websocket.Connections
	items       map[int32]*proto.Worker
	clearChan   chan int32
	ctx         context.Context
}

func (m *Manager) loopClearWorker() {
	for {
		select {
		case <-m.workers.ctx.Done():
			return
		case id, isClose := <-m.workers.clearChan:
			if !isClose {
				return
			}
			m.workers.Lock()
			delete(m.workers.items, id)
			m.workers.Unlock()
			// todo push to all
			if err := m.updateWorkerList(); err != nil {
				klog.V(2).Info(err)
			}
		}
	}
}

func (m *Manager) handlerWorker(req *proto.Request, id int32) (res []byte, err error) {
	switch req.ServiceAPI {
	case proto.ServiceAPIWorkerRegister:
		worker := &proto.Worker{}
		if err = worker.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.workers.Lock()
		m.workers.items[id] = worker
		m.workers.Unlock()
		// todo push to all
		if err = m.updateWorkerList(); err != nil {
			klog.V(2).Info(err)
			return res, err
		}
	}
	return res, nil
}

func (m *Manager) updateWorkerList() error {
	w := &proto.WorkerList{
		Items: make([]proto.Worker, 0),
	}
	m.workers.RLock()
	for _, v := range m.workers.items {
		t := *v
		w.Items = append(w.Items, t)
	}
	m.workers.RUnlock()
	res := &proto.Response{
		ServiceAPI: proto.ServiceAPIWorkerList,
		Code:       0,
		Message:    "",
		Data:       make([][]byte, 0),
	}
	if d, err := w.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	} else {
		res.Data = append(res.Data, d)
	}
	if data, err := res.Marshal(); err != nil {
		m.gateways.connections.SendToAll(data)
		return nil
	} else {
		return err
	}
}
