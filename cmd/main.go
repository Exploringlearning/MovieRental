package main

import (
	"movierental/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	err := engine.Run("localhost:8080")
	if err != nil {
		return
	}
}
