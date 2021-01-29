package manager

import (
	"context"
	"fmt"
	"github.com/nevercase/lllidan/pkg/env"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

const (
	multiplexWaitingTimeInMS = 20
	multiplexMaxLength       = 200
)

func NewWorkerHub(ctx context.Context, hostname string, registerData []byte) *workerHub {
	wh := &workerHub{
		hostname:     hostname,
		workers:      make(map[string]*worker, 0),
		ctx:          ctx,
		readChan:     make(chan []byte, 4096),
		registerData: registerData,
	}
	go wh.packageMessage()
	return wh
}

type workerHub struct {
	sync.RWMutex
	hostname     string
	workers      map[string]*worker
	ctx          context.Context
	readChan     chan []byte
	registerData []byte
}

func address(ip string, port int32) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

func (wh *workerHub) packageMessage() {
	for {
		select {
		case <-wh.ctx.Done():
			return
		case o, isClose := <-wh.readChan:
			if !isClose {
				return
			}
			req := &proto.Request{
				ServiceAPI: proto.ServiceAPIMultiplex,
			}
			res := make([][]byte, multiplexMaxLength)
			index := 0
			res[index] = o
			after := time.After(time.Millisecond * multiplexWaitingTimeInMS)
			timeout := false
			for {
				select {
				case <-wh.ctx.Done():
					return
				case <-after:
					timeout = true
					break
				case appendMsg, isClose := <-wh.readChan:
					if !isClose {
						break
					}
					index++
					res[index] = appendMsg
				}
				if timeout {
					break
				}
				if index+1 >= multiplexMaxLength {
					break
				}
			}
			req.Data = res[:index+1]
			klog.Info("req:", *req)
			data, err := req.Marshal()
			if err != nil {
				klog.V(2).Info(err)
				continue
			}
			wh.SendToAll(data)
		}
	}
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
	opt.RegisterFunc(func() error {
		return opt.Send(wh.registerData)
	})
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
			if err := opt.Send(in); err != nil {
				klog.V(2).Infof(
					"hostname:%s address:%s path:%s err:%v",
					opt.Hostname, opt.Address, opt.Path, err)
			}
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
		case in, isClose := <-w.option.Read():
			if !isClose {
				return
			}
			handleChan <- in
		}
	}
}

func GetConnectToWorkerRequest(hostname string, listenPort int32) (data []byte, err error) {
	podIp, err := env.GetPodIP()
	if err != nil {
		return data, err
	}
	gw := &proto.Gateway{
		Hostname: hostname,
		Ip:       podIp,
		Port:     listenPort,
		NodePort: 0,
	}
	if data, err = gw.Marshal(); err != nil {
		klog.V(2).Info(err)
		return data, err
	}
	req := &proto.Request{
		ServiceAPI: proto.ServiceAPIGatewayConnect,
		Data:       make([][]byte, 0),
	}
	req.Data = append(req.Data, data)
	if data, err = req.Marshal(); err != nil {
		klog.V(2).Info(err)
		return data, err
	}
	return data, nil
}
