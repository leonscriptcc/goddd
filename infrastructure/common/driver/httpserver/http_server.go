package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/leonscriptcc/goddd/infrastructure/gconfig"
	"go.uber.org/zap"
	"time"
)

// Server http服务端
type Server struct {
	engine *gin.Engine
}

// NewServer 创建http server
func NewServer() *Server {
	gin.SetMode("release")
	engine := gin.New()
	engine.Use(gin.Recovery(), zapLogger())
	return &Server{engine: engine}
}

// Start 开启http server服务
func (s *Server) Start() {
	glog.Info("http-server start listen!")
	// 开始监听
	//err := s.engine.RunTLS(
	//	":"+gconfig.Parameters.HttpServerConfig.ListeningPort,
	//	gconfig.Parameters.HttpServerConfig.CertFilePath,
	//	gconfig.Parameters.HttpServerConfig.KeyFilePath,
	//)
	err := s.engine.Run(":" + gconfig.Parameters.HttpServer.ListeningPort)
	if err != nil {
		glog.Error("http start fail:" + err.Error())
	}
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

// zapLogger 接收gin框架默认的日志
func zapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		glog.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
