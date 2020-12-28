package gateway

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
	"net/url"
	"sync"
	"time"
)

type ProxyOption struct {
	Url          url.URL
	IsMaintained bool
	Id           int32
}

type Proxy interface {
	Id() int32
	Close()
	HandlerRead() <-chan []byte
	HandlerWrite(data []byte) error
}

// NewProxy return the Proxy interface
func NewProxy(ctx context.Context, opt *ProxyOption) (wp Proxy, err error) {
	var hostname string
	if hostname, err = GetHostName(); err != nil {
		return nil, err
	}
	var c *websocket.Conn
	c, _, err = websocket.DefaultDialer.Dial(opt.Url.String(), nil)
	if err != nil {
		return nil, err
	}
	subCtx, cancel := context.WithCancel(ctx)
	p := &proxy{
		option:    opt,
		hostname:  hostname,
		readChan:  make(chan []byte, 4096),
		writeChan: make(chan []byte, 4096),
		conn:      c,
		status:    true,
		ctx:       subCtx,
		cancel:    cancel,
	}
	go p.readPump()
	go p.writePump()
	return p, nil
}

type proxy struct {
	option    *ProxyOption
	hostname  string
	readChan  chan []byte
	writeChan chan []byte
	conn      *websocket.Conn
	once      sync.Once
	status    bool
	ctx       context.Context
	cancel    context.CancelFunc
}

func (p *proxy) Id() int32 {
	return p.option.Id
}

func (p *proxy) Close() {
	p.once.Do(func() {
		p.status = false
		p.cancel()
		if err := p.conn.Close(); err != nil {
			klog.V(2).Info(err)
		}
	})
}

func (p *proxy) HandlerRead() <-chan []byte {
	return p.readChan
}

const ProxyWriteTimeout = "err: proxy write timeout 1s"
const ProxyClose = "err: proxy has been closed"

func (p *proxy) HandlerWrite(data []byte) error {
	if !p.status {
		return fmt.Errorf(ProxyClose)
	}
	timeout := time.After(time.Second * 1)
	select {
	case <-timeout:
		return fmt.Errorf(ProxyWriteTimeout)
	case p.writeChan <- data:
		return nil
	}
}

func (p *proxy) readPump() {
	defer p.Close()
	for {
		_, message, err := p.conn.ReadMessage()
		//klog.V(3).Infof("messageType: %d message: %s err:%v\n", messageType, string(message), err)
		if err != nil {
			klog.V(2).Info(err)
			return
		}
		p.readChan <- message
	}
}

func (p *proxy) writePump() {
	defer p.Close()
	for {
		select {
		case msg, isClose := <-p.writeChan:
			if !isClose {
				return
			}
			if err := p.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				klog.V(2).Info(err)
				return
			}
		case <-p.ctx.Done():
			return
		}
	}
}
