package services

import (
	"movierental/internal/dto"
	"movierental/internal/repository"
)

type Movie interface {
	Get() ([]dto.Movie, error)
	GetMoviesByFilter(string, string, string) ([]dto.Movie, error)
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

func (movie *movie) GetMoviesByFilter(genre string, year string, actor string) ([]dto.Movie, error) {
	movies, err := movie.rentalRepository.GetMoviesByFilter(genre, year, actor)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

//func fetchMovies() ([]dto.Movie, error) {
//
//	var movies []dto.Movie
//
//	resp, err := http.Get("http://www.omdbapi.com/?apikey=633cf963&s=spiderman")
//	if err != nil {
//		log.Print("Not able to get the HTTP response ", err)
//	}
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	serializeJsonToPageDto(body, &movies)
//
//	log.Println("Final list:", movies)
//	return movies, nil
//}
//
//func serializeJsonToPageDto(body []byte, movies *[]dto.Movie) {
//	if err := json.Unmarshal(body, &movies); err != nil {
//		log.Fatalln(err)
//	}
//
//}
