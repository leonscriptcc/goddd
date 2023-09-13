package ghttp

import (
	"github.com/gin-gonic/gin"
	"github.com/leonscriptcc/goddd/infrastructure/common/driver/httpserver"
)

type HttpService struct {
	demo *Demo

	httpServer *httpserver.Server
}

func NewService(demo *Demo, httpServer *httpserver.Server) *HttpService {
	return &HttpService{demo: demo, httpServer: httpServer}
}

// Start 启动http api监听
func (h *HttpService) Start() {
	// 注册路由
	register(h.httpServer.GetEngine(), h.demo)

	// 启动http服务
	h.httpServer.Start()
}

// register http路由注册
func register(engine *gin.Engine, demo *Demo) {
	engine.GET("/hello/world", demo.HelloWorld)
}
