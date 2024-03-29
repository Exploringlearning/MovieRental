package handler

import (
	"movierental/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie interface {
	Get(context *gin.Context)
	GetMoviesByFilter(context *gin.Context)
}
type movie struct {
	rentalService services.Movie
}

func NewMovie(rentalService services.Movie) Movie {
	return &movie{rentalService: rentalService}
}

func (movie *movie) Get(context *gin.Context) {
	if id := context.Param("id"); id != "" {
		movieID, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "movie id should be an integer"})
			return
		}
		rentalMovie, err := movie.rentalService.Get(movieID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, rentalMovie)
		return
	}
	movies, err := movie.rentalService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, movies)

}

func (movie *movie) GetMoviesByFilter(context *gin.Context) {
	genre := context.Query("genre")
	actor := context.Query("actor")
	year := context.Query("year")

	movies, err := movie.rentalService.GetMoviesByFilter(genre, year, actor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, movies)
}
