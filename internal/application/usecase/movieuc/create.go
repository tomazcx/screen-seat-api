package movieuc

import (
	"errors"
	"fmt"
	"time"

	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

var (
	ErrCreateMovieCategoryNotRegistered = errors.New("Category not registered")
)

type CreateMovieInput struct {
	Title       string
	Description string
	Duration    int
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	Categories  []string
}

type CreateMovieOutput struct {
	Movie     entity.Movie `json:"movie"`
	CreatedAt time.Time    `json:"created_at"`
}

type CreateMovieUsecase struct {
	movieRepository    repository.IMovieRepository
	categoryRepository repository.ICategoryRepository
}

func (uc *CreateMovieUsecase) Execute(input CreateMovieInput) (*CreateMovieOutput, error) {

	for _, category := range input.Categories {
		categoryExists, error := uc.categoryRepository.Exists(category)

		if error != nil {
			return nil, error
		}

		if !categoryExists {
			return nil, fmt.Errorf("%s: %w", category, ErrCreateMovieCategoryNotRegistered)
		}
	}

	movie, err := entity.NewMovie(input.Title, input.Description, input.Duration, "", input.AgeRating, input.Categories, input.StartDate, input.EndDate)

	if err != nil {
		return nil, err
	}

	err = uc.movieRepository.Create(movie)

	if err != nil {
		return nil, err
	}

	return &CreateMovieOutput{
		Movie:     *movie,
		CreatedAt: time.Now(),
	}, nil
}

func NewCreateMovieUseCase(
	movieRepository repository.IMovieRepository,
	categoryRepository repository.ICategoryRepository,
) *CreateMovieUsecase {
	return &CreateMovieUsecase{
		movieRepository:    movieRepository,
		categoryRepository: categoryRepository,
	}
}
