package services

import (
	"movierental/internal/dto"
	"movierental/internal/repository"
)

type Movie interface {
	Get() ([]dto.Movie, error)
}

type movie struct {
	rentalRepository repository.Movie
}

func NewMovie(rentalRepository repository.Movie) Movie {
	return &movie{rentalRepository: rentalRepository}
}

func (movie *movie) Get() ([]dto.Movie, error) {
	movies, err := movie.rentalRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return movies, nil
}
