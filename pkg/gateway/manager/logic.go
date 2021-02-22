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

func (m *Manager) registerGateway() error {
	data, err := m.registerGatewayRequest(int32(m.conf.Gateway.KubernetesService.Port))
	if err != nil {
		return err
	}
	return m.logic.option.Send(data)
}

func (m *Manager) registerGatewayRequest(listenPort int32) (data []byte, err error) {
	podIp, err := env.GetPodIP()
	if err != nil {
		return data, err
	}
	gw := &proto.Gateway{
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
		ServiceAPI: proto.ServiceAPIGatewayRegister,
		Data:       make([][]byte, 0),
	}
	req.Data = append(req.Data, data)
	if data, err = req.Marshal(); err != nil {
		klog.V(2).Info(err)
		return data, err
	}
	return data, nil
}

func (m *Manager) loopLogicMessage(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, isClose := <-m.logic.readChan:
			if !isClose {
				return
			}
			if err := m.handleLogicMessage(msg); err != nil {
				klog.V(2).Info(err)
			}
		}
	}
}

func (m *Manager) handleLogicMessage(in []byte) error {
	req := &proto.Request{}
	if err := req.Unmarshal(in); err != nil {
		klog.V(2).Info(err)
		return err
	}
	switch req.ServiceAPI {
	case proto.ServiceAPIWorkerList:
		w := &proto.WorkerList{}
		if err := w.Unmarshal(req.Data[0]); err != nil {
			klog.V(2).Info(err)
			return err
		}
		m.workers.update(w)
	}
	return nil
}
