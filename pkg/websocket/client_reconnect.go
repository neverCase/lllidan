package websocket

import (
	"k8s.io/klog/v2"
	"time"
)

func NewClientWithReconnect(opt *Option) {
	defer opt.Cancel()
	for {
		switch opt.Status {
		case OptionClosed:
			return
		case OptionActive, OptionInActive:
			client, err := NewClient(opt)
			if err != nil {
				klog.V(2).Info(err)
				time.Sleep(time.Second * 5)
				continue
			}
			opt.ChangeStatus(OptionActive)
			select {
			case <-opt.Done():
				return
			case <-client.ctx.Done():
				// reconnect
			}
		}
	}
}
