package manager

import (
	"context"
	"fmt"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"sync"
)

func NewWorkerHub(ctx context.Context, hostname string) *workerHub {
	wh := &workerHub{
		hostname:  hostname,
		workers:   make(map[string]*worker, 0),
		ctx:       ctx,
		readChan:  make(chan []byte, 4096),
		writeChan: make(chan []byte, 4096),
	}
	return wh
}

type workerHub struct {
	sync.RWMutex
	hostname  string
	workers   map[string]*worker
	ctx       context.Context
	readChan  chan []byte
	writeChan chan []byte
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
	wh.Lock()
	defer wh.Unlock()
	for k, v := range t {
		a := v
		if _, ok := wh.workers[k]; !ok {
			wh.workers[k] = wh.newWorker(&a)
		}
	}
	// cancel offline
	for k, v := range wh.workers {
		if _, ok := t[k]; !ok {
			// remove
			v.option.Cancel()
			delete(wh.workers, k)
		}
	}
}

func (wh *workerHub) newWorker(w *proto.Worker) *worker {
	opt := websocket.NewOption(
		context.Background(),
		wh.hostname,
		address(w.Ip, w.Port),
		proto.RouterGateway)
	worker := &worker{
		worker: w,
		option: opt,
	}
	go worker.readPump(wh.readChan)
	go websocket.NewClientWithReconnect(opt)
	return worker
}

func (wh *workerHub) SendToAll(in []byte) {
	wh.RLock()
	defer wh.RUnlock()
	var wg sync.WaitGroup
	wg.Add(len(wh.workers))
	for _, v := range wh.workers {
		go func(opt *websocket.Option) {
			opt.Send(in)
			wg.Done()
		}(v.option)
	}
	wg.Wait()
}

type worker struct {
	worker *proto.Worker
	option *websocket.Option
}

func (w *worker) readPump(handleChan chan<- []byte) {
	defer w.option.Cancel()
	for {
		select {
		case <-w.option.Done():
			return
		case in, isClose := <-w.option.ReadHandlerChan:
			if !isClose {
				return
			}
			handleChan <- in
		}
	}
}
