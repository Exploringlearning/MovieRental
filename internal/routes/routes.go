package routes

import (
	"movierental/internal/db"
	"movierental/internal/handler"
	"movierental/internal/repository"
	"movierental/internal/services"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	RegisterRoutes(engine)
	return engine
}

func RegisterRoutes(engine *gin.Engine) {
	dbConnection := db.SetUp()
	movieRepository := repository.NewMovie(dbConnection)
	movieService := services.NewMovie(movieRepository)
	movieHandler := handler.NewMovie(movieService)

	group := engine.Group("/movierental")
	{
		group.GET("/movies", movieHandler.Get)
		group.GET("/movies/filter", movieHandler.GetMoviesByFilter)
		group.GET("/movie/:id", movieHandler.Get)
	}
}
