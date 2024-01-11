package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMovieTitleIsRequired   = errors.New("Title is required.")
	ErrMovieInvalidDuration   = errors.New("Invalid duration. Must be a number between 5 and 300.")
	ErrMovieInvalidRate       = errors.New("Invalid rate. Must be a valid MPA film rating.")
	ErrMovieInvalidCategories = errors.New("Invalid category list. Must contain at least one category.")
	ErrMovieInvalidDates      = errors.New("Invalid dates. End date must be at least two days later than the start date.")
)

var Rates = []string{"G", "PG", "PG-13", "R", "NC-17"}

type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Poster      string    `json:"poster"`
	Duration    int       `json:"duration"`
	AgeRating   string    `json:"age_rating"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Categories  []string  `json:"categories"`
}

func (m *Movie) DurationHourString() (string, error) {
	if m.Duration < 5 || m.Duration > 300 {
		return "", ErrMovieInvalidDuration
	}
	hours := m.Duration / 60
	minutes := m.Duration % 60
	var minutesStr string
	if minutes < 10 {
		minutesStr = "0" + strconv.Itoa(minutes)
	} else {
		minutesStr = strconv.Itoa(minutes)
	}

	return fmt.Sprintf("%dh%s", hours, minutesStr), nil
}

func (m *Movie) IsMovieShowing() bool {
	now := time.Now()

	return m.EndDate.After(now)
}

func NewMovie(
	title string,
	desc string,
	duration int,
	poster string,
	rate string,
	categories []string,
	startDate time.Time,
	endDate time.Time,
) (*Movie, error) {

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

	dateDiff := endDate.Sub(startDate)
	if dateDiff < time.Hour*48 {
		return nil, ErrMovieInvalidDates
	}

	return &Movie{
		ID:          uuid.New().String(),
		Title:       title,
		Description: desc,
		Poster:      poster,
		Duration:    duration,
		AgeRating:   rate,
		Categories:  categories,
		StartDate: startDate,
		EndDate: endDate,
	}, nil
}

func (m *Movie) Validate() error {
	if len(m.Title) < 1 {
		return ErrMovieTitleIsRequired
	}

	if m.Duration < 5 || m.Duration > 300 {
		return ErrMovieInvalidDuration
	}

	if !validateRate(m.AgeRating) {
		return ErrMovieInvalidRate
	}

	if len(m.Categories) == 0 {
		return ErrMovieInvalidCategories
	}

	dateDiff := m.EndDate.Sub(m.StartDate)
	if dateDiff < time.Hour*48 {
		return ErrMovieInvalidDates
	}

	return nil
}

func validateRate(rate string) bool {
	for _, r := range Rates {
		if r == rate {
			return true
		}
	}
	return false
}
