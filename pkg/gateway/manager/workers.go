package manager

import (
	"context"
	"fmt"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"sync"
)

func NewWorkerHub() *workerHub {
	wh := &workerHub{
		workers: make(map[string]*worker, 0),
	}
	return wh
}

type workerHub struct {
	sync.RWMutex
	workers map[string]*worker
}

func address(ip string, port int32) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

func (wh *workerHub) update(in *proto.WorkerList) {
	t := make(map[string]proto.Worker, 0)
	for _, v := range in.Items {
		t[address(v.Ip, v.Port)] = v
	}
	// new online

	// cancel offline
}

func (wh *workerHub) newWorker(w *proto.Worker) {

}

type worker struct {
	worker *proto.Worker
	option *websocket.Option
	client *websocket.Client
	ctx    context.Context
	cancel context.CancelFunc
}
