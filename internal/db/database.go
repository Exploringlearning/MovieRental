package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func SetUp() *sql.DB {
	dbConn := CreateConnection()
	Migrate(dbConn)
	return dbConn
}

func CreateConnection() *sql.DB {
	dbConfig := NewDatabaseConfig()
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
		dbConfig.User,
		dbConfig.Password,
	)

	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}

	log.Println("Successfully connected to database!")
	return dbConn
}

func Migrate(dbConn *sql.DB) error {
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Println("Driver not instantiated ", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", NewDatabaseConfig().Database, driver)
	if err != nil {
		log.Fatalln("Migration with Database not instantiated ", err)
		return err
	}
	if err := m.Up(); err != nil {
		log.Println("Error occurred while migrating the database", err.Error())
		return err
	}
	log.Print("Migration completed Successfully!")
	return nil
}
