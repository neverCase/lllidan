package api

import (
	"context"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket"
	"k8s.io/klog"
	"net/url"
)

type Client struct {
	option   *websocket.Option
	readChan chan []byte
}

func GatewayUrl(host string) url.URL {
	return url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   proto.RouterGatewayApi,
	}
}

func Init(ctx context.Context, u url.URL, hostname string) *Client {
	opt := websocket.NewOption(
		ctx,
		hostname,
		u.Host,
		u.Path)
	l := &Client{
		option:   opt,
		readChan: make(chan []byte, 4096),
	}
	go websocket.NewClientWithReconnect(l.option)
	go l.readPump(l.readChan)
	return l
}

func (c *Client) Option() *websocket.Option {
	return c.option
}

func (c *Client) readPump(handleChan chan<- []byte) {
	defer c.option.Cancel()
	for {
		select {
		case <-c.option.Done():
			return
		case msg, isClose := <-c.option.Read():
			if !isClose {
				return
			}
			klog.Info("Api read msg:", msg)
		}
	}
}
