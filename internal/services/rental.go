package services

type Movie interface {
	Get() (string, error)
}
type movie struct {
}

func NewMovie() Movie {
	return &movie{}
}
func (movie *movie) Get() (string, error) {
	return "hello World", nil
}
