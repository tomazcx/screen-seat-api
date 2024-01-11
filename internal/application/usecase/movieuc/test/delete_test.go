package test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
)

func TestDeleteMovie_Success(t *testing.T){	
	input := uuid.New().String()

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("Exists", input).Return(true, (error)(nil)).Once()
	movieRepo.On("Delete", input).Return((error)(nil)).Once()

	deleteMovieUseCase := movieuc.NewDeleteMovieUseCase(movieRepo)
	err := deleteMovieUseCase.Execute(input)

	assert.Nil(t, err)
	movieRepo.AssertExpectations(t)
}

func TestDeleteMovie_Movie_Not_Found(t *testing.T){	
	input := uuid.New().String()

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("Exists", input).Return(false, (error)(nil)).Once()

	deleteMovieUseCase := movieuc.NewDeleteMovieUseCase(movieRepo)
	err := deleteMovieUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, movieuc.ErrDeleteMovieMovieNotFound))
	movieRepo.AssertExpectations(t)
	movieRepo.AssertNotCalled(t, "Delete")
}
