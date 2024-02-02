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
	service := services.NewMovie(&mockRepo)
	handler := NewMovie(service)
	mockRepo.On("GetAll").Return(make([]dto.Movie, 0), nil)
    url := "/movie"
	response := getMovie(t, url, handler.Get)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestShouldReturnAllListOfMovies(t *testing.T) {

	mockRepo := mocks.Movie{}
	service := services.NewMovie(&mockRepo)
	handler := NewMovie(service)
	movies := []dto.Movie{
		{
			ID:       1,
			Title:    "The Wolf of Wall Street",
			Year:     "2013",
			Rated:    "R",
			Released: "25 Dec 2013",
			Runtime:  "180 min",
			Genre:    "Biography, Comedy, Crime",
			Director: "Martin Scorsese",
			Writer:   "Terence Winter, Jordan Belfort",
			Actors:   "Leonardo DiCaprio, Jonah Hill, Margot Robbie",
			Plot:     "In the early 1990s, Jordan Belfort teamed with his partner Donny Azoff and started brokerage firm Stratton Oakmont...",
			Language: "English, French",
			Country:  "United States",
			Awards:   "Nominated for 5 Oscars. 37 wins & 179 nominations total",
			Poster:   "https://m.media-amazon.com/images/M/MV5BMjIxMjgxNTk0MF5BMl5BanBnXkFtZTgwNjIyOTg2MDE@._V1_SX300.jpg",
			Ratings: []dto.Rating{
				{
					Source: "Internet Movie Database",
					Value:  "8.2/10",
				},
			},
			Metascore:  "75",
			ImdbRating: "8.2",
			ImdbVotes:  "1,545,718",
			ImdbID:     "tt0993846",
			Type:       "movie",
			DVD:        "12 Dec 2015",
			BoxOffice:  "$116,900,694",
			Production: "N/A",
			Website:    "N/A",
			Response:   "True",
		},
		{
			ID:       2,
			Title:    "Interstellar",
			Year:     "2014",
			Rated:    "PG-13",
			Released: "07 Nov 2014",
			Runtime:  "169 min",
			Genre:    "Adventure, Drama, Sci-Fi",
			Director: "Christopher Nolan",
			Writer:   "Jonathan Nolan, Christopher Nolan",
			Actors:   "Matthew McConaughey, Anne Hathaway, Jessica Chastain",
			Plot:     "Earth's future has been riddled by disasters, famines, and droughts. There is only one way to ensure mankind's survival: Interstellar travel. A newly discovered wormhole in the far reaches of our solar system allows a team of astronauts to go where no man has gone before, a planet that may have the right environment to sustain human life.",
			Language: "English",
			Country:  "United States, United Kingdom, Canada",
			Awards:   "Won 1 Oscar. 44 wins & 148 nominations total",
			Poster:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg",
			Ratings: []dto.Rating{
				{Source: "Internet Movie Database", Value: "8.7/10"},
				{Source: "Rotten Tomatoes", Value: "73%"},
				{Source: "Metacritic", Value: "74/100"},
			},
			Metascore:  "74",
			ImdbRating: "8.7",
			ImdbVotes:  "2,036,452",
			ImdbID:     "tt0816692",
			Type:       "movie",
			DVD:        "24 May 2016",
			BoxOffice:  "$188,020,017",
			Production: "N/A",
			Website:    "N/A",
			Response:   "True",
		},
	}

	mockRepo.On("GetAll").Return(movies, nil)
url := "/movies"
	response := getMovie(t, url, handler.Get)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertNumberOfCalls(t, "GetAll", 1)
}

func TestShouldReturnMoviesFilterByGenre(t *testing.T) {

	engine := gin.Default()
	
	mockRepo := mocks.Movie{}
	service := services.NewMovie(&mockRepo)
	handler := NewMovie(service)
	engine.GET("movie/filter", handler.GetMoviesByFilter)


	movies := []dto.Movie{
		{
			ID:       1,
			Title:    "The Wolf of Wall Street",
			Year:     "2013",
			Rated:    "R",
			Released: "25 Dec 2013",
			Runtime:  "180 min",
			Genre:    "Biography, Comedy, Crime",
			Director: "Martin Scorsese",
			Writer:   "Terence Winter, Jordan Belfort",
			Actors:   "Leonardo DiCaprio, Jonah Hill, Margot Robbie",
			Plot:     "In the early 1990s, Jordan Belfort teamed with his partner Donny Azoff and started brokerage firm Stratton Oakmont...",
			Language: "English, French",
			Country:  "United States",
			Awards:   "Nominated for 5 Oscars. 37 wins & 179 nominations total",
			Poster:   "https://m.media-amazon.com/images/M/MV5BMjIxMjgxNTk0MF5BMl5BanBnXkFtZTgwNjIyOTg2MDE@._V1_SX300.jpg",
			Ratings: []dto.Rating{
				{
					Source: "Internet Movie Database",
					Value:  "8.2/10",
				},
			},
			Metascore:  "75",
			ImdbRating: "8.2",
			ImdbVotes:  "1,545,718",
			ImdbID:     "tt0993846",
			Type:       "movie",
			DVD:        "12 Dec 2015",
			BoxOffice:  "$116,900,694",
			Production: "N/A",
			Website:    "N/A",
			Response:   "True",
		},
		{
			ID:       2,
			Title:    "Interstellar",
			Year:     "2014",
			Rated:    "PG-13",
			Released: "07 Nov 2014",
			Runtime:  "169 min",
			Genre:    "Adventure, Drama, Sci-Fi",
			Director: "Christopher Nolan",
			Writer:   "Jonathan Nolan, Christopher Nolan",
			Actors:   "Matthew McConaughey, Anne Hathaway, Jessica Chastain",
			Plot:     "Earth's future has been riddled by disasters, famines, and droughts. There is only one way to ensure mankind's survival: Interstellar travel. A newly discovered wormhole in the far reaches of our solar system allows a team of astronauts to go where no man has gone before, a planet that may have the right environment to sustain human life.",
			Language: "English",
			Country:  "United States, United Kingdom, Canada",
			Awards:   "Won 1 Oscar. 44 wins & 148 nominations total",
			Poster:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg",
			Ratings: []dto.Rating{
				{Source: "Internet Movie Database", Value: "8.7/10"},
				{Source: "Rotten Tomatoes", Value: "73%"},
				{Source: "Metacritic", Value: "74/100"},
			},
			Metascore:  "74",
			ImdbRating: "8.7",
			ImdbVotes:  "2,036,452",
			ImdbID:     "tt0816692",
			Type:       "movie",
			DVD:        "24 May 2016",
			BoxOffice:  "$188,020,017",
			Production: "N/A",
			Website:    "N/A",
			Response:   "True",
		},
	}

	mockRepo.On("GetMoviesByFilter","action").Return(movies, nil)
    //url := "/movie/filter?genre=action"
	request, err := http.NewRequest(http.MethodGet, "/movie/filter?genre=action", nil)
	require.NoError(t, err)
	//response := getMovie(t,url,handler.GetMoviesByFilter)

	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	var responseBody []dto.Movie
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertNumberOfCalls(t, "GetMoviesByFilter", 1)
}

func getMovie(t *testing.T,url string, handler func(gin *gin.Context) ) *httptest.ResponseRecorder {
	engine := gin.Default()
	engine.GET(url, handler)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)
	return response
}
