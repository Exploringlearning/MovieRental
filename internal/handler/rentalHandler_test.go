package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"movierental/internal/dto"
	"movierental/internal/repository/mocks"
	"movierental/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldReturnStatusOK(t *testing.T) {

	mockRepo := mocks.Movie{}
	mockRepo.On("Get").Return(make([]dto.Movie, 0))

	response := getMovie(t, &mockRepo)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestShouldReturnListOfMoviesWhenSpidermanKeywordIsGiven(t *testing.T) {

	mockRepo := mocks.Movie{}
	movies := []dto.Movie{
		{
			Title: "Italian Spiderman",
			Year: "2007",
			imdbID: "tt2705436",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BZWQxMjcwNjItZjI0ZC00ZTc4LWIwMzItM2Q0YTZhNzI3NzdlXkEyXkFqcGdeQXVyMTA0MTM5NjI2._V1_SX300.jpg"
		  },
		  {
			Title: "Superman, Spiderman or Batman",
			Year: "2011",
			imdbID: "tt2084949",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BMjQ4MzcxNDU3N15BMl5BanBnXkFtZTgwOTE1MzMxNzE@._V1_SX300.jpg"
		  },
		  {
			Title: "Spiderman",
			Year: "1990",
			imdbID: "tt0100669",
			Type: "movie",
			Poster: "N/A"
		  },
		  {
			Title: "The Amazing Spiderman 2 Webb Cut",
			Year: "2021",
			imdbID: "tt18351128",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BYzYzZDViNWYtNWViMS00NDMxLThlN2YtZjFkOWMwODkzNzhiXkEyXkFqcGdeQXVyMTUwMzM4NzU0._V1_SX300.jpg"
		  },
		  {
			Title: "Spiderman",
			Year: "2010",
			imdbID: "tt1785572",
			Type: "movie",
			Poster: "N/A"
		  },
		  {
			Title: "Spiderman in Cannes",
			Year: "2016",
			imdbID: "tt5978586",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BZDlmMGQwYmItNTNmOS00OTNkLTkxNTYtNDM3ZWVlMWUyZDIzXkEyXkFqcGdeQXVyMTA5Mzk5Mw@@._V1_SX300.jpg"
		  },
		  {
			Title: "Spiderman the Verse",
			Year: "2019–",
			imdbID: "tt12122034",
			Type: "series",
			Poster: "https://m.media-amazon.com/images/M/MV5BNjA2NmZhOGEtZTQ5OS00MDI0LTg4N2UtYTRmOTllM2I2NDlhXkEyXkFqcGdeQXVyNTU4OTE5Nzc@._V1_SX300.jpg"
		  },
		  {
			Title: "Spiderman and Grandma",
			Year: "2009",
			imdbID: "tt1433184",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BMjE3Mzg0MjAxMl5BMl5BanBnXkFtZTcwNjIyODg5Mg@@._V1_SX300.jpg"
		  },
		  {
			Title: "Fighting, Flying and Driving: The Stunts of Spiderman 3",
			Year: "2007",
			imdbID: "tt1132238",
			Type: "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BNTI3NDE1ZmEtMTRiMS00YTY4LTk0OGItNjY4YmI0MDM4OGM4XkEyXkFqcGdeQXVyODE2NDgwMzM@._V1_SX300.jpg"
		  },
		  {
			Title: "Amazing Spiderman Syndrome",
			Year: "2012",
			imdbID: "tt2586634",
			Type: "movie",
			Poster: "N/A"
		  },
	}

	mockRepo.On("Get").Return(movies)

	response := getMovie(t, &mockRepo)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertNumberOfCalls(t, "Get", 1)
}

func getMovie(t *testing.T, mockRepo *mocks.Movie) *httptest.ResponseRecorder {
	engine := gin.Default()
	service := services.NewMovie(mockRepo)
	handler := NewMovie(service)
	engine.GET("/movie", handler.Get)

	request, err := http.NewRequest(http.MethodGet, "/movie", nil)
	require.NoError(t, err)

	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)
	return response
}
