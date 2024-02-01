package repository

import "movierental/internal/dto"

type Movie interface {
	Get() []dto.Movie
}

type movie struct {
}

func NewMovie() Movie {
	return &movie{}
}

func (m *movie) Get() []dto.Movie {
	return make([]dto.Movie, 0)
}
