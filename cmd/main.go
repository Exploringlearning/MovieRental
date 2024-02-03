package main

import (
	"github.com/gin-gonic/gin"
	"movierental/internal/routes"
)

func main() {
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
