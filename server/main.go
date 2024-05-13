package main

import (
	"github.com/gin-gonic/gin"
	"vuego/server/internal/routes"
)

func main() {
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	engine.Run(":8082")
}
