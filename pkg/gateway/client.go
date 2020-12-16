package gateway

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/logic"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"net/url"
	"sync"
	"time"
)

func NewClientWithRecover(c *config.Config, stop <-chan struct{}, handler chan<- []byte) {
	for {
		client, err := NewClient(c, handler)
		if err != nil {
			klog.V(2).Info(err)
			time.Sleep(time.Second * 5)
			continue
			//return
		}
		select {
		case <-client.done():
		case <-stop:
			return
		}
	}
}

func NewClient(conf *config.Config, handler chan<- []byte) (*Client, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", conf.Logic.KubernetesService.Name, conf.Logic.KubernetesService.Port),
		Path:   proto.RouterGateway,
	}
	a, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	var hostname string
	if hostname, err = GetHostName(); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{
		hostname:    hostname,
		conf:        conf,
		conn:        a,
		writeChan:   make(chan []byte, 1024),
		handlerChan: handler,
		ctx:         ctx,
		cancel:      cancel,
	}
	if err = c.registerGateway(); err != nil {
		return nil, err
	}
	go c.readPump()
	go c.writePump()
	go c.ping()
	return c, nil
}

type Client struct {
	hostname    string
	conf        *config.Config
	conn        *websocket.Conn
	writeChan   chan []byte
	handlerChan chan<- []byte
	once        sync.Once
	ctx         context.Context
	cancel      context.CancelFunc
}

func (c *Client) done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Client) close() {
	c.once.Do(func() {
		c.cancel()
	})
}

func (c *Client) ping() {
	defer c.close()
	tick := time.NewTicker(time.Second * time.Duration(logic.WebsocketConnectionTimeout/2))
	defer tick.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-tick.C:
			req := &proto.Request{
				ServiceAPI: proto.ServiceAPIPing,
				Data:       make([][]byte, 0),
			}
			data, err := req.Marshal()
			if err != nil {
				klog.V(2).Info(err)
				return
			}
			c.writeChan <- data
		}
	}
}

func (c *Client) readPump() {
	defer c.close()
	for {
		_, message, err := c.conn.ReadMessage()
		//klog.V(3).Infof("messageType: %d message: %s err:%v\n", messageType, string(message), err)
		if err != nil {
			klog.V(2).Info(err)
			return
		}
		c.handlerChan <- message
	}
}

func (c *Client) writePump() {
	defer c.close()
	for {
		select {
		case msg, isClose := <-c.writeChan:
			if !isClose {
				return
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				klog.V(2).Info(err)
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) registerGateway() (err error) {
	gw := &proto.Gateway{
		Hostname: c.hostname,
		Ip:       c.conf.Gateway.KubernetesService.Name,
		Port:     int32(c.conf.Gateway.KubernetesService.Port),
		NodePort: int32(c.conf.Gateway.KubernetesService.NodePort),
	}
	var data []byte
	if data, err = gw.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	}
	req := &proto.Request{
		ServiceAPI: proto.ServiceAPIGatewayRegister,
		Data:       make([][]byte, 0),
	}
	req.Data = append(req.Data, data)
	if data, err = req.Marshal(); err != nil {
		klog.V(2).Info(err)
		return err
	}
	c.writeChan <- data
	return nil
}
