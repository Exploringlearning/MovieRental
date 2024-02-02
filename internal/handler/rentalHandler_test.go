package handler

import (
	"movierental/internal/dto"
	"movierental/internal/repository/mocks"
	"movierental/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnStatusOK(t *testing.T) {

	mockRepo := mocks.Movie{}
	mockRepo.On("GetAll").Return(make([]dto.Movie, 0), nil)

	response := getMovie(t, &mockRepo)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestShouldReturnListOfMovies(t *testing.T) {

	mockRepo := mocks.Movie{}
	movies := []dto.Movie{
		{
			Title:  "Italian Spiderman",
			Year:   "2007",
			ImdbId: "tt2705436",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BZWQxMjcwNjItZjI0ZC00ZTc4LWIwMzItM2Q0YTZhNzI3NzdlXkEyXkFqcGdeQXVyMTA0MTM5NjI2._V1_SX300.jpg",
		},
		{
			Title:  "Superman, Spiderman or Batman",
			Year:   "2011",
			ImdbId: "tt2084949",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BMjQ4MzcxNDU3N15BMl5BanBnXkFtZTgwOTE1MzMxNzE@._V1_SX300.jpg",
		},
		{
			Title:  "Spiderman the Verse",
			Year:   "2019–",
			ImdbId: "tt12122034",
			Type:   "series",
			Poster: "https://m.media-amazon.com/images/M/MV5BNjA2NmZhOGEtZTQ5OS00MDI0LTg4N2UtYTRmOTllM2I2NDlhXkEyXkFqcGdeQXVyNTU4OTE5Nzc@._V1_SX300.jpg",
		},
	}

	mockRepo.On("GetAll").Return(movies, nil)

	response := getMovie(t, &mockRepo)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertNumberOfCalls(t, "GetAll", 1)
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
