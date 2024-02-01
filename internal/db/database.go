package db

import (
	"database/sql"
	"fmt"
	
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func SetUp() *sql.DB {
	dbConn := CreateConnection()
	Migrate(dbConn)
	return dbConn
}

func CreateConnection() *sql.DB {
	dbConfig := NewConfig()
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
		dbConfig.DBUser,
		dbConfig.DBPassword,
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
		log.Print("Driver not instantiated ", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", NewConfig().DBName, driver)
	if err != nil {
		log.Print("Migration with Database not instantiated ", err)
		return err
	}
	m.Up()
	log.Print("Migration Successfull")
	return nil
}
