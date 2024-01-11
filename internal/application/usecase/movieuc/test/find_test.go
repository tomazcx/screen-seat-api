package test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestFindMovieUseCase_Success(t *testing.T){
	input := uuid.New().String()
	movie := &entity.Movie{ID: input, Title: "Test movie"}

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("FindById", input).Return(movie, (error)(nil)).Once()

	findMovieUseCase := movieuc.NewFindMovieUseCase(movieRepo)

	output, err := findMovieUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, output, movie)
	movieRepo.AssertExpectations(t)
}

func TestFindMovieUseCase_Movie_Not_Found(t *testing.T){
	input := uuid.New().String()
	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("FindById", input).Return((*entity.Movie)(nil), errors.New("database error")).Once()

	findMovieUseCase := movieuc.NewFindMovieUseCase(movieRepo)

	output, err := findMovieUseCase.Execute(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, movieuc.ErrFindMovieMovieNotFound))
	movieRepo.AssertExpectations(t)
}


