package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	email string `json:"email"`
	password string `json:"password"`
}

type Register struct {
	firstname string
	lastname string
	email string
	password string
}

func main() {
  r := gin.Default()

	v1 := r.Group("/v1") 
	{
		v1.POST("/register", func(ctx *gin.Context) {
			var register Register

			if err := ctx.ShouldBind(&register); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		})

		v1.POST("/login", func(ctx *gin.Context) {
			var json Login

			if err := ctx.ShouldBind(&json); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		})

		v1.POST("/logout", func(ctx *gin.Context) {
		})

		endpoints := Endpoints{group: v1}

		endpoints.initializeUserHandlers(useTestMiddlware(), useAuthenticationMiddleware())
		endpoints.initializeHealthHandlers(useTestMiddlware())

		c := make(chan string, 1)
		endpoints.initTestHandlers(c)
	}

	r.Run(":8081")
}

type Endpoints struct {
	group *gin.RouterGroup
}

func (e Endpoints) initializeUserHandlers(middlewares ...gin.HandlerFunc) {
	users := e.group.Group("/users")
	{
		users.Use(middlewares...)
		users.GET("/", func(ctx *gin.Context) {})
		users.POST("/", func(ctx *gin.Context) {})

		users.GET("/:user", func(ctx *gin.Context) {
			user := ctx.Param("user")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello " + user,
			})
			return
		})
		users.POST("/:user", func(ctx *gin.Context) {
			user := ctx.Param("user")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello " + user,
			})
			return
		})
	}
}

func (e Endpoints) initializeHealthHandlers(middlewares ...gin.HandlerFunc) {
	health := e.group.Group("/health")
	{
		health.Use(middlewares...)
		health.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "I'm healthy",
			})
			return
		})
	}
}
func (e Endpoints) initAuthHandlers(middlewares ...gin.HandlerFunc) {
	auth := e.group.Group("/auth")
	{
		auth.Use(middlewares...)
		auth.POST("/login", func(ctx *gin.Context) {
		})
		auth.POST("/register", func(ctx *gin.Context) {
		})
		auth.DELETE("/logout", func(ctx *gin.Context) {
		})
	}
}

func (e Endpoints) initTestHandlers(c chan string, middlwares ...gin.HandlerFunc) {
	test := e.group.Group("/test")
	{
		test.Use(middlwares...)
		test.GET("/receive", func(ctx *gin.Context) {
			select {
				case msg := <-c: 
					ctx.JSON(http.StatusOK, gin.H{
						"msg": msg,
					})
				default: 
					ctx.JSON(http.StatusOK, gin.H{
						"msg": "Nothing in the buffer for you",
					})
			}
			return
		})
		test.POST("/send/:msg", func(ctx *gin.Context) {
			msg := ctx.Param("msg")
			c <- msg
			ctx.Status(http.StatusOK)
			return
		})
	}
}

func useAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if (rand.Float64() >= 0.5) {
			fmt.Println("----> Authentication Middleware chose pass")
			ctx.Next()
		} else {
			fmt.Println("----> Authentication Middleware chose error")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func useTestMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Test Middleware invoked")
	}
}
