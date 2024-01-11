package movieuc

import (
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

type FindManyMoviesInput struct {
	Page     int
	Limit    int
	SortBy   string
	Title    string
	Rate     string
	Category string
}

type FindManyMoviesUseCase struct {
	movieRepository repository.IMovieRepository
}

func (uc *FindManyMoviesUseCase) Execute(input FindManyMoviesInput) ([]entity.Movie, error) {
	return uc.movieRepository.FindMany(input.Page, input.Limit, input.SortBy, input.Title, input.Rate, input.Category)
}

func NewFindManyMoviesUseCase(movieRepo repository.IMovieRepository) *FindManyMoviesUseCase {
	return &FindManyMoviesUseCase{
		movieRepository: movieRepo,
	}
}
