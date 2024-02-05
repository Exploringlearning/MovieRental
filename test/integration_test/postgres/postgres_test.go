package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"movierental/internal/dto"
	"movierental/internal/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	engine *gin.Engine
)

func TestMain(m *testing.M) {
	testContainer, dbConn := Setup()
	exitCode := m.Run()
	Teardown(testContainer, dbConn)
	os.Exit(exitCode)
}

func Setup() (testcontainers.Container, *sql.DB) {
	ctx := context.Background()
	dbContainer, dbConn, err := Postgres(ctx)
	if err != nil {
		log.Fatalln("Error while setting up test database", err.Error())
	}

	engine = gin.Default()
	routes.RegisterRoutes(engine)

	return dbContainer, dbConn
}

func Teardown(container testcontainers.Container, dbConn *sql.DB) {
	ctx := context.Background()
	err := container.Terminate(ctx)
	if err != nil {
		log.Fatalln("Error while terminating the container")
	}
	err = dbConn.Close()
	if err != nil {
		log.Fatalln("Error while closing the connection")
	}
}

func TestMovieRentalHandler(t *testing.T) {

	t.Run("TestShouldReturnAllMovies", func(t *testing.T) {
		movies := []dto.Movie{
			{
				ID:         1,
				Title:      "The Wolf of Wall Street",
				Year:       "2013",
				Rated:      "R",
				Released:   "25 Dec 2013",
				Runtime:    "180 min",
				Genre:      "Biography, Comedy, Crime",
				Director:   "Martin Scorsese",
				Writer:     "Terence Winter, Jordan Belfort",
				Actors:     "Leonardo DiCaprio, Jonah Hill, Margot Robbie",
				Plot:       "In the early 1990s, Jordan Belfort teamed with his partner Donny Azoff and started brokerage firm Stratton Oakmont...",
				Language:   "English, French",
				Country:    "United States",
				Awards:     "Nominated for 5 Oscars. 37 wins & 179 nominations total",
				Poster:     "https://m.media-amazon.com/images/M/MV5BMjIxMjgxNTk0MF5BMl5BanBnXkFtZTgwNjIyOTg2MDE@._V1_SX300.jpg",
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
				ID:         2,
				Title:      "Interstellar",
				Year:       "2014",
				Rated:      "PG-13",
				Released:   "07 Nov 2014",
				Runtime:    "169 min",
				Genre:      "Adventure, Drama, Sci-Fi",
				Director:   "Christopher Nolan",
				Writer:     "Jonathan Nolan, Christopher Nolan",
				Actors:     "Matthew McConaughey, Anne Hathaway, Jessica Chastain",
				Plot:       "Earth's future has been riddled by disasters, famines, and droughts. There is only one way to ensure mankind's survival: Interstellar travel. A newly discovered wormhole in the far reaches of our solar system allows a team of astronauts to go where no man has gone before, a planet that may have the right environment to sustain human life.",
				Language:   "English",
				Country:    "United States, United Kingdom, Canada",
				Awards:     "Won 1 Oscar. 44 wins & 148 nominations total",
				Poster:     "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg",
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
			{
				ID:         3,
				Title:      "Oppenheimer",
				Year:       "2023",
				Rated:      "R",
				Released:   "21 Jul 2023",
				Runtime:    "180 min",
				Genre:      "Biography, Drama, History",
				Director:   "Christopher Nolan",
				Writer:     "Christopher Nolan, Kai Bird, Martin Sherwin",
				Actors:     "Cillian Murphy, Emily Blunt, Matt Damon",
				Plot:       "The story of American scientist J. Robert Oppenheimer and his role in the development of the atomic bomb.",
				Language:   "English, German, Italian",
				Country:    "United States, United Kingdom",
				Awards:     "102 wins & 231 nominations",
				Poster:     "https://m.media-amazon.com/images/M/MV5BMDBmYTZjNjUtN2M1MS00MTQ2LTk2ODgtNzc2M2QyZGE5NTVjXkEyXkFqcGdeQXVyNzAwMjU2MTY@._V1_SX300.jpg",
				Metascore:  "88",
				ImdbRating: "8.4",
				ImdbVotes:  "577,705",
				ImdbID:     "tt15398776",
				Type:       "movie",
				DVD:        "21 Nov 2023",
				BoxOffice:  "$326,102,235",
				Production: "N/A",
				Website:    "N/A",
				Response:   "True",
			},
		}
		url := "http://localhost:8080/movierental/movies"
		response := getResponse(t, url)

		var responseBody []dto.Movie
		err := json.NewDecoder(response.Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, movies, responseBody)
		assert.NoError(t, err)
	})
	t.Run("TestShouldReturnMovieForGivenId", func(t *testing.T) {
		expectedMovie := dto.Movie{
			ID:         1,
			Title:      "The Wolf of Wall Street",
			Year:       "2013",
			Rated:      "R",
			Released:   "25 Dec 2013",
			Runtime:    "180 min",
			Genre:      "Biography, Comedy, Crime",
			Director:   "Martin Scorsese",
			Writer:     "Terence Winter, Jordan Belfort",
			Actors:     "Leonardo DiCaprio, Jonah Hill, Margot Robbie",
			Plot:       "In the early 1990s, Jordan Belfort teamed with his partner Donny Azoff and started brokerage firm Stratton Oakmont...",
			Language:   "English, French",
			Country:    "United States",
			Awards:     "Nominated for 5 Oscars. 37 wins & 179 nominations total",
			Poster:     "https://m.media-amazon.com/images/M/MV5BMjIxMjgxNTk0MF5BMl5BanBnXkFtZTgwNjIyOTg2MDE@._V1_SX300.jpg",
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
		url := "http://localhost:8080/movierental/movie/1"
		response := getResponse(t, url)

		var responseBody dto.Movie
		err := json.NewDecoder(response.Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedMovie, responseBody)
		assert.NoError(t, err)
	})
	//t.Run("TestShouldFilterMoviesByGenre", func(t *testing.T) {
	//
	//})

}

func getResponse(t *testing.T, url string) *httptest.ResponseRecorder {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)
	return response
}
