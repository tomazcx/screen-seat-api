package test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestUpdateMovieUseCase_Success(t *testing.T) {
	movieId := uuid.New().String()
	movieRepo := &mocks.MovieRepositoryMock{}
	categoryRepo := &mocks.CategoryRepositoryMock{}

	movieToUpdate := &entity.Movie{
		ID:          movieId,
		Title:       "TestTitle",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories:  []string{"Action", "Adventure"},
		StartDate:   time.Now().AddDate(0, 0, 2),
		EndDate:     time.Now().AddDate(0, 0, 7),
	}

	movieRepo.On("Exists", movieId).Return(true, (error)(nil))
	movieRepo.On("FindById", movieId).Return(movieToUpdate, (error)(nil))
	movieRepo.On("Update", movieToUpdate).Return((error)(nil))
	categoryRepo.On("Exists", "Action").Return(true, (error)(nil))

	useCase := movieuc.NewUpdateMovieUseCase(movieRepo, categoryRepo)

	input := movieuc.UpdateMovieInput{
		Title:       "Test movie updated",
		Description: "Test description updated",
		Duration:    120,
		AgeRating:   "G",
		Categories:  []string{"Action"},
		StartDate:   time.Now().AddDate(0, 0, 1),
		EndDate:     time.Now().AddDate(0, 0, 9),
	}

	output, err := useCase.Execute(movieId, input)

	assert.Nil(t, err)
	assert.Equal(t, output.ID, movieId)
	assert.Equal(t, output.Title, input.Title)
	assert.Equal(t, output.Description, input.Description)
	assert.Equal(t, output.Duration, input.Duration)
	assert.Equal(t, output.AgeRating, input.AgeRating)
	assert.Equal(t, output.Categories, input.Categories)
	movieRepo.AssertCalled(t, "Update", mock.Anything)
	movieRepo.AssertCalled(t, "Exists", movieId)
	categoryRepo.AssertCalled(t, "Exists", "Action")
}

func TestUpdateMovieUseCase_Movie_Not_Found(t *testing.T) {
	movieId := uuid.New().String()
	movieRepo := &mocks.MovieRepositoryMock{}
	categoryRepo := &mocks.CategoryRepositoryMock{}

	movieRepo.On("Exists", movieId).Return(false, (error)(nil))

	useCase := movieuc.NewUpdateMovieUseCase(movieRepo, categoryRepo)

	input := movieuc.UpdateMovieInput{
		Title:       "Test movie",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories:  []string{"Action", "Adventure"},
		StartDate:   time.Now().AddDate(0, 0, 2),
		EndDate:     time.Now().AddDate(0, 0, 7),
	}

	output, err := useCase.Execute(movieId, input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.True(t, errors.Is(err, movieuc.ErrUpdateMovieMovieNotFound))
	movieRepo.AssertNotCalled(t, "Update")
	movieRepo.AssertCalled(t, "Exists", movieId)
	categoryRepo.AssertNotCalled(t, "Exists")
	categoryRepo.AssertNotCalled(t, "FindById")
}

func TestUpdateMovieUseCase_Category_Not_Found(t *testing.T) {
	movieId := uuid.New().String()
	movieRepo := &mocks.MovieRepositoryMock{}
	categoryRepo := &mocks.CategoryRepositoryMock{}

	movieRepo.On("Exists", movieId).Return(true, (error)(nil))
	categoryRepo.On("Exists", "Action").Return(false, (error)(nil))

	useCase := movieuc.NewUpdateMovieUseCase(movieRepo, categoryRepo)

	input := movieuc.UpdateMovieInput{
		Title:       "Test movie",
		Description: "Test description",
		Duration:    90,
		AgeRating:   "NC-17",
		Categories:  []string{"Action", "Adventure"},
		StartDate:   time.Now().AddDate(0, 0, 2),
		EndDate:     time.Now().AddDate(0, 0, 7),
	}

	output, err := useCase.Execute(movieId, input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.True(t, errors.Is(err, movieuc.ErrUpdateMovieCategoryRegistered))
	movieRepo.AssertNotCalled(t, "Update")
	movieRepo.AssertCalled(t, "Exists", movieId)
	categoryRepo.AssertCalled(t, "Exists", "Action")
	categoryRepo.AssertNotCalled(t, "FindById")
}
