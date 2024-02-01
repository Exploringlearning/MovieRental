package repository

import (
	"database/sql"
	"movierental/internal/dto"
)

type Movie interface {
	Get() []dto.Movie
}

type movie struct {
	*sql.DB
}

func NewMovie(db *sql.DB) Movie {
	return &movie{db}
}

func (m *movie) Get() []dto.Movie {
	return make([]dto.Movie, 0)
}
