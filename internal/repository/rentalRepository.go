package repository

import (
	"database/sql"
	"log"
	"movierental/internal/dto"
)

type Movie interface {
	GetAll() ([]dto.Movie, error)
	GetMoviesByFilter(string) ([]dto.Movie, error)
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

	movies, err := m.scanMovie(rows)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *movie) GetMoviesByFilter(genre string)([]dto.Movie, error){
	
	rows, err := m.DB.Query("SELECT * from movies where genre = ?",genre)

	if err != nil {
		log.Println("Error occurred while fetching movies from database", err.Error())
		return nil, err
	}
	defer rows.Close()

	movies, err := m.scanMovie(rows)
	if err != nil {
		return nil, err
	}	

	return movies,nil
}

func (m *movie) scanMovie(rows *sql.Rows) ([]dto.Movie, error) {
	var movies []dto.Movie

	for rows.Next() {
		var movie dto.Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			&movie.Rated,
			&movie.Released,
			&movie.Runtime,
			&movie.Genre,
			&movie.Director,
			&movie.Writer,
			&movie.Actors,
			&movie.Plot,
			&movie.Language,
			&movie.Country,
			&movie.Awards,
			&movie.Poster,
			&movie.Metascore,
			&movie.ImdbRating,
			&movie.ImdbVotes,
			&movie.ImdbID,
			&movie.Type,
			&movie.DVD,
			&movie.BoxOffice,
			&movie.Production,
			&movie.Website,
			&movie.Response,
		); err != nil {
			log.Println("Error occurred while serializing movies from result", err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
