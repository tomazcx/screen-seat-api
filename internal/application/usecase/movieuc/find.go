package movieuc

import (
	"errors"
	"fmt"

	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

var (
	ErrFindMovieMovieNotFound = errors.New("Movie could not be found")
)

type FindMovieUseCase struct {
	movieRepository repository.IMovieRepository
}

func (uc *FindMovieUseCase) Execute(id string) (*entity.Movie, error) {
	movie, err := uc.movieRepository.FindById(id)

	if err != nil {
		return nil, fmt.Errorf("%w:%w", ErrFindMovieMovieNotFound, err)
	}

	return movie, nil
}

func NewFindMovieUseCase(movieRepository repository.IMovieRepository) *FindMovieUseCase {
	return &FindMovieUseCase{
		movieRepository: movieRepository,
	}
}
