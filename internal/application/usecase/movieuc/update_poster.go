package movieuc

import (
	"errors"

	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
)

var (
	ErrUpdatePosterMovieNotFound = errors.New("Movie not found")
)

type UpdateMoviePosterInput struct {
	MovieId   string
	PosterUrl string
}

type UpdateMoviePosterUseCase struct {
	movieRepository repository.IMovieRepository
}

func (uc *UpdateMoviePosterUseCase) Execute(input UpdateMoviePosterInput) error {

	movieExists, err := uc.movieRepository.Exists(input.MovieId)

	if err != nil {
		return err
	}

	if !movieExists {
		return ErrUpdatePosterMovieNotFound
	}

	err = uc.movieRepository.UpdatePoster(input.MovieId, input.PosterUrl)

	return err
}

func NewUpdateMoviePosterUseCase(movieRepository repository.IMovieRepository) *UpdateMoviePosterUseCase {
	return &UpdateMoviePosterUseCase{
		movieRepository: movieRepository,
	}
}
