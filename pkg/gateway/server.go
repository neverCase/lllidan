package gateway

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/gateway/manager"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	c           *config.Config
	server      *http.Server
	connections *proxyConnections
	manager     *manager.Manager
	ctx         context.Context
}

func Init(c *config.Config) *Server {
	m, err := manager.NewManger(context.Background(), c)
	if err != nil {
		klog.Fatal(err)
	}
	s := &Server{
		c:           c,
		connections: NewConnections(context.Background()),
		manager:     m,
	}
	router := gin.New()
	router.Use(cors.Default())
	for _, v := range c.Gateway.Routes {
		router.GET(v.Path, s.handler)
	}
	router.GET(proto.RouterGatewayApi, s.wsHandlerApi)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", c.Gateway.ListenPort),
		Handler: router,
	}
	s.server = server
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				klog.Info("Server closed under request")
			} else {
				klog.V(2).Info("Server closed unexpected err:", err)
			}
		}
	}()
	return s
}

func (s *Server) handler(c *gin.Context) {
	var u url.URL
	isExist := false
	for _, v := range s.c.Gateway.Routes {
		if len(v.Parameters) == 0 {
			continue
		}
		params := make([]string, 0)
		for _, p := range v.Parameters {
			if !p.IsDynamic {
				params = append(params, p.DefaultValue)
			} else {
				val := c.Param(p.Key)
				if p.IsRequired && val == "" {
					continue
				}
				params = append(params, val)
			}
		}
		u = url.URL{
			Scheme: v.Scheme,
			Host:   v.Host,
			Path:   fmt.Sprintf("/%s", strings.Join(params, "/")),
		}
		isExist = true
		break
	}
	if !isExist {
		return
	}
	s.connections.handler(c.Writer, c.Request, u)
}

func (s *Server) wsHandlerApi(c *gin.Context) {
	s.manager.Handler(c.Writer, c.Request, proto.RouterGatewayApi)
}
