package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestCreateMovie_Success(t *testing.T){
	testTitle := "testMovie"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	categories := []string{"Action"}

	movie, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.Nil(t, err)
	assert.NotNil(t, movie.ID)
	assert.Equal(t, movie.Title, testTitle)
	assert.Equal(t, movie.Description, testDescription)
	assert.Equal(t, movie.Duration, testDuration)
	assert.Equal(t, movie.AgeRating, testRating)
	assert.Equal(t, movie.Categories, categories)
}

func TestCreateMovie_Invalid_Title(t *testing.T){
	testTitle := ""
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	categories := []string{"Action"}

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieTitleIsRequired)
}

func TestCreateMovie_Invalid_Duration_Below_5(t *testing.T){
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 4
	testRating := "G"
	categories := []string{"Action"}

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)
}

func TestCreateMovie_Invalid_Duration_Bigger_Than_300(t *testing.T){
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 301
	testRating := "G"
	categories := []string{"Action"}

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)
}

func TestCreateMovie_Invalid_Test_Rating(t *testing.T){
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "X"
	categories := []string{"Action"}

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidRate)
}

func TestCreateMovie_Invalid_Categories(t *testing.T){
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	categories := []string{}

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, testRating, categories)

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidCategories)
}

func TestMovieDurationHourString_Success(t *testing.T) {
	movie, err := entity.NewMovie("Test", "", 90, "G", []string{"Action"})

	assert.Nil(t, err)

	expected := "1h30"
	durationHourString, err := movie.DurationHourString()

	assert.Nil(t, err)
	assert.Equal(t, expected, durationHourString)
}

func TestMovieDurationHourString_Fail(t *testing.T) {
	movie, err := entity.NewMovie("Test", "", 90, "G", []string{"Action"})

	assert.Nil(t, err)

	movie.Duration = 0
	_, err := movie.DurationHourString()

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)
}

