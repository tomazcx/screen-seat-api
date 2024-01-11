package movieuc

import (
	"errors"
	"fmt"
	"time"

	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

var (
	ErrUpdateMovieCategoryRegistered = errors.New("Category not registered")
	ErrUpdateMovieMovieNotFound      = errors.New("Movie not found")
)

type UpdateMovieInput struct {
	Title       string
	Description string
	Duration    int
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	Categories  []string
}

type UpdateMovieUsecase struct {
	movieRepository    repository.IMovieRepository
	categoryRepository repository.ICategoryRepository
}

func (uc *UpdateMovieUsecase) Execute(movieId string, input UpdateMovieInput) (*entity.Movie, error) {

	movieExists, err := uc.movieRepository.Exists(movieId)

	if err != nil {
		return nil, err
	}

	if !movieExists {
		return nil, ErrUpdateMovieMovieNotFound
	}

	for _, category := range input.Categories {
		categoryExists, error := uc.categoryRepository.Exists(category)

		if error != nil {
			return nil, error
		}

		if !categoryExists {
			return nil, fmt.Errorf("%s: %w", category, ErrUpdateMovieCategoryRegistered)
		}
	}

	if err != nil {
		return nil, err
	}

	movie, err := uc.movieRepository.FindById(movieId)

	if err != nil {
		return nil, err
	}

	err = assignMovieValues(movie, &input)

	if err != nil {
		return nil, err
	}

	err = uc.movieRepository.Update(movie)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func assignMovieValues(movie *entity.Movie, input *UpdateMovieInput) error {
	movie.Title = input.Title
	movie.Description = input.Description
	movie.Duration = input.Duration
	movie.StartDate = input.StartDate
	movie.EndDate = input.EndDate
	movie.AgeRating = input.AgeRating
	movie.Categories = input.Categories
	return movie.Validate()
}

func NewUpdateMovieUseCase(
	movieRepository repository.IMovieRepository,
	categoryRepository repository.ICategoryRepository,
) *UpdateMovieUsecase {
	return &UpdateMovieUsecase{
		movieRepository:    movieRepository,
		categoryRepository: categoryRepository,
	}
}
