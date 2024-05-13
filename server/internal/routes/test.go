package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implTest struct {
	c chan string
}

func initTestHandler(router *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	impl := implTest{c: make(chan string, 1)}
	
	test := router.Group("/test")
	{
		test.Use(middlewares...)
		test.GET("/receive", impl.receiveMessage)
		test.POST("/send/:msg", impl.sendMessage)
	}
}

func (impl implTest) receiveMessage(ctx *gin.Context) {
	select {
		case msg := <-impl.c: 
		ctx.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		default: 
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Nothing in the buffer for you",
		})
	}
}

func (impl implTest) sendMessage(ctx *gin.Context) {
	msg := ctx.Param("msg")
	impl.c <- msg
	ctx.Status(http.StatusOK)
}

