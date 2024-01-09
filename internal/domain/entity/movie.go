package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMovieTitleIsRequired = errors.New("Title is required.")
	ErrMovieInvalidDuration = errors.New("Invalid duration. Must be a number between 5 and 300.")
	ErrMovieInvalidRate = errors.New("Invalid rate. Must be a valid MPA film rating.")
	ErrMovieInvalidCategories = errors.New("Invalid category list. Must contain at least one category.")
)

var Rates = []string{"G", "PG", "PG-13", "R", "NC-17"}

type Movie struct {
	ID          string
	Title       string
	Description string
	Duration    int
	AgeRating   string
	Categories    []string
}

func NewMovie(title string, desc string, duration int, rate string, categories []string) (*Movie, error) {
	if len(title) < 1 {
		return nil, ErrMovieTitleIsRequired
	}

	if duration < 5 || duration > 300 {
		return nil, ErrMovieInvalidDuration
	}

	if !validateRate(rate) {
		return nil, ErrMovieInvalidRate
	}

	if len(categories) == 0 {
		return nil, ErrMovieInvalidCategories
	}

	return &Movie{
		ID: uuid.New().String(),
		Title: title,
		Description: desc,
		Duration: duration,
		AgeRating: rate,
		Categories: categories,
	}, nil
}


func validateRate(rate string) bool {
	for _, r := range Rates {
		if r == rate {
			return true
		}
	}
	return false
}
