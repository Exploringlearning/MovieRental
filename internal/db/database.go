package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

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
