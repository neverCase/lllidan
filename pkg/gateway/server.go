package gateway

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nevercase/lllidan/pkg/config"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	c           *config.Config
	server      *http.Server
	connections *connections
	ctx         context.Context
}

func Init(c *config.Config) *Server {
	s := &Server{
		c:           c,
		connections: NewConnections(context.Background()),
	}
	router := gin.New()
	router.Use(cors.Default())
	for _, v := range c.Gateway.Routes {
		router.GET(v.Path, s.handler)
	}
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
		params := make([]string, len(v.Parameters))
		for k, p := range v.Parameters {
			if k == 0 {
				svc := c.Param(p.Key)
				if svc != p.DefaultValue {
					break
				}
				params[k] = p.DefaultValue
			} else {
				if p.IsDynamic {
					params[k] = c.Param(p.Key)
				} else {
					params[k] = p.DefaultValue
				}
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
