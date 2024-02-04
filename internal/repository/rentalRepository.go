package repository

import (
	"database/sql"
	"fmt"
	"log"
	"movierental/internal/dto"
)

type Movie interface {
	GetAll() ([]dto.Movie, error)
	GetMoviesByFilter(string, string, string) ([]dto.Movie, error)
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

func (m *movie) GetMoviesByFilter(genre string, year string, actor string) ([]dto.Movie, error) {

	query := getQueryForFilter(genre, year, actor)

	rows, err := m.DB.Query(query)

	if err != nil {
		log.Println("Error occurred while fetching movies from database", err.Error())
		return nil, err
	}
	defer rows.Close()

	movies, err := m.scanMovie(rows)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func getQueryForFilter(genre string, year string, actor string) string {
	query := "SELECT * FROM MOVIES WHERE"

	var conditions []string

	if genre != "" {
		conditions = append(conditions, fmt.Sprintf(" genre ILIKE '%s' ", wrapInPercentage(genre)))
	}

	if year != "" {
		conditions = append(conditions, fmt.Sprintf(" year = '%s' ", year))
	}

	if actor != "" {
		conditions = append(conditions, fmt.Sprintf(" actor ILIKE '%s' ", wrapInPercentage(actor)))
	}

	for i := 0; i < len(conditions); i++ {
		query += conditions[i]
		if i < len(conditions)-1 {
			query += "AND"
		}
	}
	return query
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

func wrapInPercentage(str string) string {
	return "%" + str + "%"
}
