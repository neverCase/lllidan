package logic

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/proto"
	"k8s.io/klog/v2"
	"net/http"
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
	router.GET(proto.RouterDashboard, s.handlerDashboard)
	router.GET(proto.RouterGateway, s.handlerGateway)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", c.Logic.ListenPort),
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

func (s *Server) handlerDashboard(c *gin.Context) {
	//s.connections.handler(c.Writer, c.Request, connTypeDashboard, )
}

func (s *Server) handlerGateway(c *gin.Context) {

}
