package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestCreateMovie_Success(t *testing.T) {
	testTitle := "testMovie"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	poster := "image.jpeg"
	categories := []string{"Action"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)

	movie, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.Nil(t, err)
	assert.NotNil(t, movie.ID)
	assert.Equal(t, movie.Title, testTitle)
	assert.Equal(t, movie.Description, testDescription)
	assert.Equal(t, movie.Duration, testDuration)
	assert.Equal(t, movie.AgeRating, testRating)
	assert.Equal(t, movie.Categories, categories)
}

func TestCreateMovie_Invalid_Title(t *testing.T) {
	testTitle := ""
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	categories := []string{"Action"}
	poster := "image.jpeg"
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieTitleIsRequired)
}

func TestCreateMovie_Invalid_Duration(t *testing.T) {
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 4
	testRating := "G"
	poster := "image.jpeg"
	categories := []string{"Action"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)

	testDuration = 301
	_, err = entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)
}
func TestCreateMovie_Invalid_Test_Rating(t *testing.T) {
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "X"
	poster := "image.jpeg"
	categories := []string{"Action"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidRate)
}

func TestCreateMovie_Invalid_Categories(t *testing.T) {
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 90
	testRating := "G"
	categories := []string{}
	startDate := time.Now()
	poster := "image.jpeg" 
	endDate := startDate.AddDate(0, 2, 0)

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidCategories)
}

func TestCreateMovie_Invalid_Date(t *testing.T) {
	testTitle := "Test title"
	testDescription := "testDescription"
	testDuration := 90
	poster := "image.jpeg" 
	testRating := "G"
	categories := []string{"Action"}
	startDate := time.Now()
	endDate := startDate

	_, err := entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDates)

	endDate = startDate.AddDate(0, 0, 1)
	_, err = entity.NewMovie(testTitle, testDescription, testDuration, poster, testRating, categories, startDate, endDate)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDates)

}

func TestIsMovieShowing(t *testing.T){
	startDate := time.Now().AddDate(0, -2, 0)
	endDate := startDate.AddDate(0, 0, 7)
	movie, err := entity.NewMovie("Test", "", 90, "poster", "G", []string{"Action"}, startDate, endDate)
	assert.Nil(t, err)

	result := movie.IsMovieShowing()
	assert.False(t, result)

	startDate = time.Now()
	endDate = startDate.AddDate(0, 0, 7)
	movie.StartDate = startDate
	movie.EndDate = endDate

	result = movie.IsMovieShowing()
	assert.True(t, result)
}

func TestMovieDurationHourString_Batch_Success(t *testing.T) {
	batchValues := []struct {
		Value    int
		Expected string
	}{
		{Value: 60, Expected: "1h00"},
		{Value: 68, Expected: "1h08"},
		{Value: 74, Expected: "1h14"},
		{Value: 82, Expected: "1h22"},
		{Value: 97, Expected: "1h37"},
		{Value: 103, Expected: "1h43"},
		{Value: 116, Expected: "1h56"},
		{Value: 127, Expected: "2h07"},
		{Value: 134, Expected: "2h14"},
		{Value: 145, Expected: "2h25"},
		{Value: 158, Expected: "2h38"},
		{Value: 162, Expected: "2h42"},
		{Value: 173, Expected: "2h53"},
		{Value: 181, Expected: "3h01"},
		{Value: 189, Expected: "3h09"},
		{Value: 196, Expected: "3h16"},
		{Value: 207, Expected: "3h27"},
		{Value: 218, Expected: "3h38"},
		{Value: 229, Expected: "3h49"},
		{Value: 240, Expected: "4h00"},
	}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)

	for _, item := range batchValues {
		movie, err := entity.NewMovie("Test", "", item.Value, "image.jpeg", "G", []string{"Action"}, startDate, endDate) 
		assert.Nil(t, err)

		durationHourString, err := movie.DurationHourString()
		assert.Nil(t, err)
		assert.Equal(t, item.Expected, durationHourString)
	}
}

func TestMovieDurationHourString_Fail(t *testing.T) {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)
	movie, err := entity.NewMovie("Test", "", 90, "image.jpeg", "G", []string{"Action"}, startDate, endDate)
	assert.Nil(t, err)

	movie.Duration = 0
	_, err = movie.DurationHourString()

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrMovieInvalidDuration)
}
