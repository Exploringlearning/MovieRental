package repository

import (
	"database/sql"
	"log"
	"movierental/internal/dto"
)

type Movie interface {
	GetAll() ([]dto.Movie, error)
}

type movie struct {
	*sql.DB
}

func NewMovie(db *sql.DB) Movie {
	return &movie{db}
}

func (m *movie) GetAll() ([]dto.Movie, error) {
	rows, err := m.DB.Query("SELECT * FROM Movies")
	if err != nil {
		log.Println("Error occurred while fetching all movies from database", err.Error())
		return nil, err
	}
	defer rows.Close()

	var movies []dto.Movie

	for rows.Next() {
		var movie dto.Movie
		if err = rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.ImdbId, &movie.Type, &movie.Poster); err != nil {
			log.Println("Error occurred while serializing movies from result", err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
