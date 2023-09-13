package wire

import (
	"github.com/leonscriptcc/goddd/api/ghttp"
	"github.com/leonscriptcc/goddd/application"
	"github.com/leonscriptcc/goddd/infrastructure/common/driver/httpserver"
)

// Wire 依赖注入
func Wire() *ghttp.HttpService {
	app := application.NewApp()
	api := ghttp.NewDemo(app)
	httpService := ghttp.NewService(api, httpserver.NewServer())
	return httpService
}
