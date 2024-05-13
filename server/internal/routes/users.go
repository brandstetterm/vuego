package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {}

func initUserHandler(router *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	handler := userHandler{}

	users := router.Group("/users")
	{
		users.Use(middlewares...)
		users.GET("/", handler.getUsers)
		users.POST("/", handler.createUser)
		users.GET("/:id", handler.getUser)
		users.PUT("/:id", handler.updateUser)
		users.DELETE("/:id", handler.deleteUser)
	}
}

func (_ userHandler) getUsers(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
func (_ userHandler) createUser(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
func (_ userHandler) getUser(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
func (_ userHandler) updateUser(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
func (_ userHandler) deleteUser(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
