package dto

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbId string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
