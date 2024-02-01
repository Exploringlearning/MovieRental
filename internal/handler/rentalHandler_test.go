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
			Title:  "SpiderMan",
			Year:   "1999",
			ImdbId: "3524",
			Type:   "Action",
			Poster: "https://www.poster.com/get/3524",
		},
		{
			Title:  "SpiderMan-2",
			Year:   "2010",
			ImdbId: "3534",
			Type:   "Action",
			Poster: "https://www.poster.com/get/3534",
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
