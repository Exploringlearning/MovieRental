package routes

import (
	"movierental/internal/handler"
	"movierental/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {	
	movieService := services.NewMovie()
	movieHandler := handler.NewMovie(movieService)
	engine.GET("/", movieHandler.Get)
}
