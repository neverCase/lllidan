package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
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

func (m *Manager) handlerWorker(in []byte, handler handler.Interface) (res []byte, err error) {
	req := &proto.Request{}
	if err = req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIPing:
		handler.Ping()
	case proto.ServiceAPIWorkerRegister:
		worker := &proto.Worker{}
		if err = worker.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return nil, err
		}
		m.workers.Lock()
		m.workers.items[handler.Id()] = worker
		m.workers.Unlock()
		klog.Infof("handlerWorker items:%v", m.workers.items)
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
	res := &proto.Request{
		ServiceAPI: proto.ServiceAPIWorkerList,
		Data:       make([][]byte, 0),
	}
	if d, err := w.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	} else {
		res.Data = append(res.Data, d)
	}
	klog.Info("updateWorkerList res:", *res)
	if data, err := res.Marshal(); err != nil {
		return err
	} else {
		m.gateways.connections.SendToAll(data)
		return nil
	}
}
