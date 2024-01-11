package movieuc

import (
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

type FindExhibitionMoviesUseCase struct {
	movieRepository repository.IMovieRepository
}

func (uc *FindExhibitionMoviesUseCase) Execute() ([]entity.Movie, error) {
	return uc.movieRepository.FindAllExhbition()
}

func NewFindExhibitionMoviesUseCase(movieRepository repository.IMovieRepository) *FindExhibitionMoviesUseCase {
	return &FindExhibitionMoviesUseCase{
		movieRepository: movieRepository,
	}
}
