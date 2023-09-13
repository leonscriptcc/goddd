package main

import (
	"github.com/leonscriptcc/goddd/infrastructure/gconfig"
	"github.com/leonscriptcc/goddd/infrastructure/glog"
	"github.com/leonscriptcc/goddd/wire"
	"log"
)

func init() {
	// 载入配置文件
	if err := gconfig.Load(); err != nil {
		log.Panic("load config fail:", err)
	}

	// 初始化日志服务
	if err := glog.Init(); err != nil {
		log.Panic("init log fail:", err)
	}
}

func main() {
	httpService := wire.Wire()

	// 启动http服务
	httpService.Start()
}
