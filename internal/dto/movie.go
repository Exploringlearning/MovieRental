package dto

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbId string `json:"imdbId"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}
