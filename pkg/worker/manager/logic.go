package manager

import (
	"context"
	"github.com/nevercase/lllidan/pkg/env"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"k8s.io/klog/v2"
	"net/url"
)

type logic struct {
	option   *websocket.Option
	readChan chan []byte
}

func (l *logic) readPump(handleChan chan<- []byte) {
	defer l.option.Cancel()
	for {
		select {
		case <-l.option.Done():
			return
		case in, isClose := <-l.option.Read():
			if !isClose {
				return
			}
			// todo handler
			handleChan <- in
		}
	}
}

func InitLogic(ctx context.Context, u url.URL, hostname string) *logic {
	data, err := websocket.PingData()
	if err != nil {
		klog.Fatal(err)
	}
	opt := websocket.NewOption(
		ctx,
		hostname,
		u.Host,
		u.Path,
		data)
	l := &logic{
		option:   opt,
		readChan: make(chan []byte, 4096),
	}
	go l.readPump(l.readChan)
	return l
}

func (m *Manager) registerWorker() error {
	data, err := m.registerWorkerRequest(int32(m.conf.Worker.KubernetesService.Port))
	if err != nil {
		return err
	}
	return m.logic.option.Send(data)
}

func (m *Manager) registerWorkerRequest(listenPort int32) (data []byte, err error) {
	podIp, err := env.GetPodIP()
	if err != nil {
		return data, err
	}
	gw := &proto.Worker{
		Hostname: m.hostname,
		Ip:       podIp,
		Port:     listenPort,
		NodePort: 0,
	}
	if data, err = gw.Marshal(); err != nil {
		klog.V(2).Info(err)
		return data, err
	}
	req := &proto.Request{
		ServiceAPI: proto.ServiceAPIWorkerRegister,
		Data:       make([][]byte, 0),
	}
	req.Data = append(req.Data, data)
	if data, err = req.Marshal(); err != nil {
		klog.V(2).Info(err)
		return data, err
	}
	return data, nil
}
