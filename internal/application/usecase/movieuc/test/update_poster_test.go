package test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
)

func TestUpdateMoviePoster(t *testing.T){
	movieId := uuid.New().String() 
	posterUrl := "image.jpeg"

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("Exists", movieId).Return(true, (error)(nil)).Once()
	movieRepo.On("UpdatePoster", movieId, posterUrl).Return((error)(nil)).Once()

	updateMoviePosterUseCase := movieuc.NewUpdateMoviePosterUseCase(movieRepo)

	input := movieuc.UpdateMoviePosterInput{
		MovieId: movieId,
		PosterUrl: posterUrl,
	}

	err := updateMoviePosterUseCase.Execute(input)
	
	assert.Nil(t, err)
	movieRepo.AssertExpectations(t)
}

func TestUpdateMoviePoster_Movie_Not_Found(t *testing.T){
	movieId := uuid.New().String() 
	posterUrl := "image.jpeg"

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("Exists", movieId).Return(false, (error)(nil)).Once()

	updateMoviePosterUseCase := movieuc.NewUpdateMoviePosterUseCase(movieRepo)

	input := movieuc.UpdateMoviePosterInput{
		MovieId: movieId,
		PosterUrl: posterUrl,
	}

	err := updateMoviePosterUseCase.Execute(input)
	
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, movieuc.ErrUpdatePosterMovieNotFound))
	movieRepo.AssertExpectations(t)
	movieRepo.AssertNotCalled(t, "UpdatePoster")
}
