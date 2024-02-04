package handler

import (
	"github.com/pkg/errors"
	"movierental/internal/dto"
	"movierental/internal/repository/mocks"
	"movierental/internal/services"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	engine   *gin.Engine
	mockRepo mocks.Movie
	service  services.Movie
	handler  Movie
)

func setup() {
	engine = gin.Default()
	mockRepo = mocks.Movie{}
	service = services.NewMovie(&mockRepo)
	handler = NewMovie(service)

	engine.GET("/movies", handler.Get)
	engine.GET("/movies/filter", handler.GetMoviesByFilter)
	engine.GET("/movie/:id", handler.Get)

}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestShouldReturnAllListOfMovies(t *testing.T) {
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
	response := getResponse(t, url)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertCalled(t, "GetAll")
}

func TestShouldReturnMoviesFilterByGenre(t *testing.T) {
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
	}

	genre := "crime"

	mockRepo.On("GetMoviesByFilter", genre, "", "").Return(movies, nil)

	url := "/movies/filter?genre=" + genre
	response := getResponse(t, url)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertCalled(t, "GetMoviesByFilter", genre, "", "")
}

func TestShouldReturnMoviesFilterByGenreAndYear(t *testing.T) {
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
	}

	genre := "crime"
	year := "2013"

	mockRepo.On("GetMoviesByFilter", genre, year, "").Return(movies, nil)

	url := "/movies/filter?genre=" + genre + "&year=" + year
	response := getResponse(t, url)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertCalled(t, "GetMoviesByFilter", genre, year, "")
}

func TestShouldReturnMoviesFilterByGenreAndYearAndActor(t *testing.T) {
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
	}

	genre := "crime"
	year := "2013"
	actor := "Margot Robbie"

	mockRepo.On("GetMoviesByFilter", genre, year, actor).Return(movies, nil)

	url := "/movies/filter?genre=" + genre + "&year=" + year + "&actor=" + actor
	response := getResponse(t, url)

	var responseBody []dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movies, responseBody)
	mockRepo.AssertCalled(t, "GetMoviesByFilter", genre, year, actor)
}

func TestShouldReturnMovieForAGivenMovieID(t *testing.T) {
	movieDTO := dto.Movie{
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
	}

	movieID := 1

	mockRepo.On("Get", movieID).Return(movieDTO, nil)

	url := "/movie/" + strconv.Itoa(movieID)
	response := getResponse(t, url)

	var responseBody dto.Movie
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, movieDTO, responseBody)
	mockRepo.AssertCalled(t, "Get", movieID)
}

func TestShouldReturnErrorWhenMovieNotFound(t *testing.T) {
	movieID := 0

	mockRepo.On("Get", movieID).Return(dto.Movie{}, errors.New("movie not found"))

	url := "/movie/" + strconv.Itoa(movieID)
	response := getResponse(t, url)

	var responseBody map[string]string
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	expectedResponse := map[string]string{"error": "movie not found"}

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, expectedResponse, responseBody)
	mockRepo.AssertCalled(t, "Get", movieID)
}

func TestShouldReturnErrorWhenInvalidIdIsGiven(t *testing.T) {
	movieID := "a"

	url := "/movie/" + movieID
	response := getResponse(t, url)

	var responseBody map[string]string
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)

	expectedResponse := map[string]string{"error": "movie id should be an integer"}

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, expectedResponse, responseBody)
}

func getResponse(t *testing.T, url string) *httptest.ResponseRecorder {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)
	return response
}
