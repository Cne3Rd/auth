package routers

import (
	"github.com/cne3rd/sadsec/controllers"
	"github.com/gin-gonic/gin"
)

type Control struct {
	Route *gin.Engine
	App   *controllers.Application
}

func (ctx *Control) Routes() {

	ctx.Route.GET("/register", ctx.App.Register)
	ctx.Route.GET("/login", ctx.App.Login)
	ctx.Route.POST("/register", ctx.App.Register)
	ctx.Route.POST("/login", ctx.App.Login)
	ctx.Route.GET("/logout", ctx.App.Logout)

	ctx.Route.GET("/verify_your_email/:username/:verify", ctx.App.VerifyEmail)

	ir := ctx.Route.Group("/", ctx.App.CheckSession)
	ir.GET("/", ctx.App.Index)

	ctx.Route.POST("/500", ctx.App.InternalServerError)
}
