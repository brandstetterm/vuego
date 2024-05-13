package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implAuth struct {}

func initAuthHandler(router *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	impl := implAuth{}

	auth := router.Group("/auth")
	{
		auth.Use(middlewares...)
		auth.POST("/login", impl.login)
		auth.POST("/signup", impl.register)
		auth.DELETE("/logout", impl.logout)
	}
}

func (impl implAuth) login(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
	return
}
func (impl implAuth) register(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
	return
}
func (impl implAuth) logout(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
	return
}

