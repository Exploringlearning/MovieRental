package services

import (
	"encoding/json"
	"io"
	"log"
	"movierental/internal/dto"
	"movierental/internal/repository"
	"net/http"
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
	
	var movies dto.MovieSearch

	resp, err := http.Get("http://www.omdbapi.com/?apikey=633cf963&s=spiderman")
	if err != nil {
		log.Print("Not able to get the HTTP response ", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	serializeJsonToPageDto(body, &movies)

	log.Println("Final list:", movies)
	return movies.Search, nil
}


func serializeJsonToPageDto(body []byte, movies *dto.MovieSearch) {
	if err := json.Unmarshal(body, &movies); err != nil {
		log.Fatalln(err)
	}

}