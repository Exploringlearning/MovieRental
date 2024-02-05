package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"movierental/internal/db"
)

const (
	TestDBUser     = "user"
	TestDBPassword = "password"
	TestDatabase   = "movie-rental-db-test"
)

func Postgres(context context.Context) (testcontainers.Container, *sql.DB, error) {
	container, connectionString := setupPostgresContainer()

	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("Error while pinging the database")
	}

	err = db.Migrate(dbConn, "file://../../../internal/db/migrations")
	if err != nil {
		log.Fatalln("Error while migrating the database")
	}

	return container, dbConn, nil
}

func setupPostgresContainer() (testcontainers.Container, string) {
	ctx := context.Background()

	log.Println("Started creating Test Container")
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:latest"),
		postgres.WithDatabase(TestDatabase),
		postgres.WithUsername(TestDBUser),
		postgres.WithPassword(TestDBPassword),
		testcontainers.WithWaitStrategy(
			wait.ForExposedPort(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Println("Error while creating connection string")
	}

	return container, connStr
}

func getDBConfigEnvVariable(host string, port string) []string {
	var variables []string
	variables = append(variables, fmt.Sprintf("DB_USER=%s", TestDBUser))
	variables = append(variables, fmt.Sprintf("DB_PASSWORD=%s", TestDBPassword))
	variables = append(variables, fmt.Sprintf("DB_HOST=%s", host))
	variables = append(variables, fmt.Sprintf("DB_PORT=%s", port))
	variables = append(variables, fmt.Sprintf("DB_NAME=%s", TestDatabase))
	return variables
}
