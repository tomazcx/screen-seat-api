package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func CreateMovieUseCase_Success(t *testing.T) {
	movieRepo := &mocks.MovieRepository{}
	categoryRepo := &mocks.CategoryRepository{}

	movieRepo.On("Create", mock.Anything).Return((error)(nil))
	categoryRepo.On("Exists", "Action").Return((error)(nil))
	categoryRepo.On("Exists", "Adventure").Return((error)(nil))

	useCase := movieuc.NewCreateMovieUseCase(suite.movieRepo, suite.categoryRepo)

	input := movieuc.CreateMovieInput{
		Title:       "Test movie",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories: []string{"Action", "Adventure"},
		StartDate: time.Now().AddDate(0, 0, 2),
		EndDate: time.Now().AddDate(0, 0, 7),
	}

	movie, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, movie.ID)
	assert.Equal(t, movie.Title, input.Title)
	assert.Equal(t, movie.Description, input.Description)
	assert.Equal(t, movie.Duration, input.Duration)
	assert.Equal(t, movie.AgeRating, input.AgeRating)
	assert.Equal(t, movie.Categories, input.Categories)
	movieRepo.AssertCalled(t, "Create", mock.Anything)
	categoryRepo.AssertCalled(t, "Exists", "Action")
	categoryRepo.AssertCalled(t, "Exists", "Adventure")
}
