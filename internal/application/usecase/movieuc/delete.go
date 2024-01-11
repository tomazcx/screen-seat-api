package movieuc

import (
	"errors"

	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
)

var (
	ErrDeleteMovieMovieNotFound = errors.New("Movie not found")
)

type DeleteMovieUseCase struct {
	movieRepository repository.IMovieRepository
}

func (uc *DeleteMovieUseCase) Execute(id string) error {

	movieExists, err := uc.movieRepository.Exists(id)

	if err != nil {
		return err
	}

	if !movieExists {
		return ErrDeleteMovieMovieNotFound
	}

	err = uc.movieRepository.Delete(id)

	return err
}

func NewDeleteMovieUseCase(movieRepository repository.IMovieRepository) *DeleteMovieUseCase {
	return &DeleteMovieUseCase{
		movieRepository: movieRepository,
	}
}
