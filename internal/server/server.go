package server

import (
	"context"
	"github.com/cjphaha/eDefender/common/middleware"
	"github.com/cjphaha/eDefender/internal/service"
	"github.com/cjphaha/eDefender/plugin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server interface {
	Start()
}

type server struct {
	ctx    context.Context
	cancel context.CancelFunc
	router *gin.Engine
	srv    service.Service
	c      *Config
}

func New(ctx context.Context, cancel context.CancelFunc, srv service.Service, c *Config) (Server, error) {
	// base config
	s := &server{
		ctx:    ctx,
		cancel: cancel,
		srv:    srv,
		router: gin.Default(),
		c:      c,
	}

	if s.c.IsCORS {
		log.Info("cors middleware opened")
		s.router.Use(middleware.Cors())
	}

	s.SetRouter()

	return s, nil
}

func (s *server) Start() {
	log.Info("gin http server start at ", s.c)
	s.router.Run(":" + s.c.Port)
}

func (s *server) SetRouter() {
	g := s.router.Group("/api")

	s.unAuthRouter(g)
}

func (s *server) unAuthRouter(g *gin.RouterGroup) {
	g.GET("/pluginList", func(c *gin.Context) {
		c.JSON(200, plugin.GetPlugins())
	})
	g.POST("/check", func(c *gin.Context) {
		var json plugin.Task
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := plugin.Scan(json)
		c.JSON(200, result)
	})
	g.GET("/info", func(c *gin.Context) {
		info := s.srv.GetInfo()
		c.JSON(http.StatusOK, info)
	})
}

func (s *server) Close() {
	s.cancel()
}