package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestFindManyMovies(t *testing.T){
	movies := []entity.Movie{
		{ID: uuid.New().String(), Title: "Test movie 1", AgeRating: "G", Categories: []string{"Action", "Adventure"}},
		{ID: uuid.New().String(), Title: "Test movie 22", AgeRating: "G", Categories: []string{"Animation", "Adventure"}},
		{ID: uuid.New().String(), Title: "Test movie 3", AgeRating: "G", Categories: []string{"Adventure"}},
	}

	input := movieuc.FindManyMoviesInput{
		Page: 1,
		Limit:3,
		SortBy: "created_at",
		Title: "Test",
		Rate: "G",
		Category: "Adventure",
	}

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("FindMany", input.Page, input.Limit, input.SortBy, input.Title, input.Rate, input.Category).Return(movies, (error)(nil)).Once()
	
	findMovieUseCase := movieuc.NewFindManyMoviesUseCase(movieRepo)

	output, err := findMovieUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, output, movies)
	movieRepo.AssertExpectations(t)

}
