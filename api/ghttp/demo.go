package ghttp

import (
	"github.com/gin-gonic/gin"
	"github.com/leonscriptcc/goddd/application"
)

type Demo struct {
	app *application.App
}

func NewDemo(app *application.App) *Demo {
	return &Demo{app: app}
}

func (d *Demo) HelloWorld(ctx *gin.Context) {

}
