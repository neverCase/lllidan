package worker

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/proto"
	"github.com/nevercase/lllidan/pkg/websocket/handler"
	"github.com/nevercase/lllidan/pkg/worker/manager"
	"k8s.io/klog/v2"
	"net/http"
)

type Server struct {
	c       *config.Config
	server  *http.Server
	manager *manager.Manager
	ctx     context.Context
}

func Init(c *config.Config) *Server {
	m, err := manager.NewManager(context.Background(), c)
	if err != nil {
		klog.Fatal(err)
	}
	s := &Server{
		c:       c,
		manager: m,
	}
	router := gin.New()
	router.Use(cors.Default())
	router.GET(proto.RouterGateway, s.wsHandlerGateway)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", c.Worker.ListenPort),
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

func (s *Server) GatewayMessageHandler() handler.MessageHandler {
	return s.manager.GatewayMessageHandler()
}

func (s *Server) Shutdown() {}

func (s *Server) wsHandlerGateway(c *gin.Context) {
	s.manager.Handler(c.Writer, c.Request, proto.RouterGateway)
}
