package main

import (
	"fmt"
	"vuego/server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	if err := engine.Run(":8082"); err != nil {
		fmt.Println("Error trying to run the http server")
	}
}
