package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
)

func setupPostgresContainer(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()

	log.Println("Started creating Test Container")
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:latest"),
		postgres.WithDatabase("movierental-test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForExposedPort(),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	log.Println("Connection String -> ", connStr)

	assert.NoError(t, err)

	return container, connStr
}

func TestPostgresSQLIntegration(t *testing.T) {
	container, connectionString := setupPostgresContainer(t)
	defer func(container testcontainers.Container, ctx context.Context) {
		err := container.Terminate(ctx)
		if err != nil {
			log.Println("Error while terminating the container..")
		}
	}(container, context.Background())

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Error while pinging the database")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error while closing the db connection..")
		}
	}(db)

	t.Run("TestShouldCreateMovieTable", func(t *testing.T) {
		_, err = db.Exec("CREATE TABLE movies (ID INT NOT NULL, Title VARCHAR(255), Year VARCHAR(4), Rated VARCHAR(10));")
		assert.NoError(t, err)
	})

	t.Run("TestShouldInsertValuesIntoMovieTable", func(t *testing.T) {
		_, err = db.Exec("INSERT INTO movies (ID, Title, Year, Rated) VALUES (1,'Interstellar','2010','9.0/10'), (2,'Tenet','2020','8.8/10');")
		assert.NoError(t, err)
	})

	t.Run("TestShouldSelectValuesFromMoviesTable", func(t *testing.T) {
		rows, err := db.Query("SELECT * FROM movies;")
		assert.NoError(t, err)
		defer rows.Close()

		type Movie struct {
			ID    int
			Title string
			Year  string
			Rated string
		}

		var expectedResult = []Movie{
			{
				ID:    1,
				Title: "Interstellar",
				Year:  "2010",
				Rated: "9.0/10",
			},
			{
				ID:    2,
				Title: "Tenet",
				Year:  "2020",
				Rated: "8.8/10",
			},
		}

		var actualResult []Movie

		for rows.Next() {
			var movie Movie
			err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Rated)
			assert.NoError(t, err)
			actualResult = append(actualResult, movie)
		}

		assert.ElementsMatch(t, expectedResult, actualResult)
	})
}
