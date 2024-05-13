package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implHealth struct {}

func initHealthHandler(router *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	impl := implHealth{}
	
	health := router.Group("/health")
	{
		health.Use(middlewares...)
		health.GET("/", impl.getHealthStatus)
	}
}

func (_ implHealth) getHealthStatus(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Healthy")
}
