package main

import (
	"movierental/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	//engine.GET("/", hello)
	routes.RegisterRoutes(engine)
	engine.Run("localhost:8080")
}	