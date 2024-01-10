package test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
)

func TestCreateMovieUseCase_Success(t *testing.T) {
	movieRepo := &mocks.MovieRepositoryMock{}
	categoryRepo := &mocks.CategoryRepositoryMock{}

	movieRepo.On("Create", mock.Anything).Return((error)(nil))
	categoryRepo.On("Exists", "Action").Return(true, (error)(nil))
	categoryRepo.On("Exists", "Adventure").Return(true, (error)(nil))

	useCase := movieuc.NewCreateMovieUseCase(movieRepo, categoryRepo)

	input := movieuc.CreateMovieInput{
		Title:       "Test movie",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories: []string{"Action", "Adventure"},
		StartDate: time.Now().AddDate(0, 0, 2),
		EndDate: time.Now().AddDate(0, 0, 7),
	}

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, output.Movie.ID)
	assert.Equal(t, output.Movie.Title, input.Title)
	assert.Equal(t, output.Movie.Description, input.Description)
	assert.Equal(t, output.Movie.Duration, input.Duration)
	assert.Equal(t, output.Movie.AgeRating, input.AgeRating)
	assert.Equal(t, output.Movie.Categories, input.Categories)
	movieRepo.AssertCalled(t, "Create", mock.Anything)
	categoryRepo.AssertCalled(t, "Exists", "Action")
	categoryRepo.AssertCalled(t, "Exists", "Adventure")
}

func TestCreateMovieUseCase_Category_Not_Registered(t *testing.T){	
	movieRepo := &mocks.MovieRepositoryMock{}
	categoryRepo := &mocks.CategoryRepositoryMock{}

	movieRepo.On("Create", mock.Anything).Return((error)(nil))
	categoryRepo.On("Exists", "Action").Return(false, (error)(nil))

	useCase := movieuc.NewCreateMovieUseCase(movieRepo, categoryRepo)

	input := movieuc.CreateMovieInput{
		Title:       "Test movie",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories: []string{"Action"},
		StartDate: time.Now().AddDate(0, 0, 2),
		EndDate: time.Now().AddDate(0, 0, 7),
	}

	_, err := useCase.Execute(input)
	
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, movieuc.ErrCreateMovieCategoryNotRegistered))
}
