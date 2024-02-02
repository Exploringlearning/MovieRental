package routes

import (
	"movierental/internal/db"
	"movierental/internal/handler"
	"movierental/internal/repository"
	"movierental/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	dbConnection := db.SetUp()
	movieRepository := repository.NewMovie(dbConnection)
	movieService := services.NewMovie(movieRepository)
	movieHandler := handler.NewMovie(movieService)
	engine.GET("/movierental", movieHandler.Get)
	engine.GET("/movierental/filter", movieHandler.GetMoviesByFilter)
}
