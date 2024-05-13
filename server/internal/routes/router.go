package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	{
		initUserHandler(api)
		initHealthHandler(api)
		initTestHandler(api)
	}
}
