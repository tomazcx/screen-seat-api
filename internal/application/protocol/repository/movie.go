package repository

import "github.com/tomazcx/screen-seat-api/internal/domain/entity"

type IMovieRepository interface {
	Exists(id string) (bool, error)
	FindById(id string) (*entity.Movie, error)
	FindAllExhbition() ([]entity.Movie, error)
	FindMany(page int, limit int, sort string, title string, rate string, category string) ([]entity.Movie, error)
	Create(*entity.Movie) error
	UpdatePoster(id string, posterUrl string) error
	Update(*entity.Movie) error
	Delete(id string) error
}
