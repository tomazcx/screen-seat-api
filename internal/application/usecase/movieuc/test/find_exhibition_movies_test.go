package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/screen-seat-api/internal/application/usecase/movieuc"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestFindExhibitionMovies(t *testing.T){
	now := time.Now()
	expected := []entity.Movie{
		{ID: uuid.New().String(), Title: "Movie one", StartDate: now.AddDate(0, 0, -7), EndDate: now.AddDate(0, 0, 7) },
		{ID: uuid.New().String(), Title: "Movie one", StartDate: now.AddDate(0, 0, -10), EndDate: now.AddDate(0, 0, 5) },
		{ID: uuid.New().String(), Title: "Movie one", StartDate: now.AddDate(0, 0, -2), EndDate: now.AddDate(0, 0, 20) },
	}

	movieRepo := &mocks.MovieRepositoryMock{}
	movieRepo.On("FindAllExhbition").Return(expected, (error)(nil))

	findMoviesInExhibitionUseCase := movieuc.NewFindExhibitionMoviesUseCase(movieRepo)
	output, err := findMoviesInExhibitionUseCase.Execute()

	assert.Nil(t, err)
	assert.Equal(t, output, expected)

	for _, movie := range output {
		assert.True(t, movie.EndDate.After(now))
	}
}
