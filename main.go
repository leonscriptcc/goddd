package main

import (
	"github.com/leonscriptcc/goddd/infrastructure/logs"
	"log"
)

func init() {
	// 初始化日志服务
	if err := logs.Init(); err != nil {
		log.Panic("logs init fail:", err)
	}
}

func main() {
}
