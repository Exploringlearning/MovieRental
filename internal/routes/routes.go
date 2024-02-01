package routes

import (
	"movierental/internal/handler"
	"movierental/internal/repository"
	"movierental/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	movieRepository := repository.NewMovie()
	movieService := services.NewMovie(movieRepository)
	movieHandler := handler.NewMovie(movieService)
	engine.GET("/", movieHandler.Get)
}
