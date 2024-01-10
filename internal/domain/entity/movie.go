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
	ErrMovieSessionRequired   = errors.New("At least one subtitled or a dubbed movie session is required.")
	ErrMovieInvalidSession    = errors.New("One of the defined movie sessions is not valid. All days should contain at least one movie session with hour between 1pm and 10h30pm and a room.")
)

var Rates = []string{"G", "PG", "PG-13", "R", "NC-17"}

type MovieSession struct {
	Hour float32
	Room string
}

type MovieWeekdaySessions map[time.Weekday][]MovieSession

type Movie struct {
	ID          string
	Title       string
	Description string
	Duration    int
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	DubSessions MovieWeekdaySessions
	SubSessions MovieWeekdaySessions
	Categories  []string
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
	rate string,
	categories []string,
	startDate time.Time,
	endDate time.Time,
	dubSessions MovieWeekdaySessions,
	subSessions MovieWeekdaySessions,
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

	if len(dubSessions) == 0 && len(subSessions) == 0 {
		return nil, ErrMovieSessionRequired
	}

	if !validateMovieSessions(dubSessions) {
		return nil, ErrMovieInvalidSession
	}

	if !validateMovieSessions(subSessions) {
		return nil, ErrMovieInvalidSession
	}

	return &Movie{
		ID:          uuid.New().String(),
		Title:       title,
		Description: desc,
		Duration:    duration,
		AgeRating:   rate,
		Categories:  categories,
	}, nil
}

func validateMovieSessions(movieSessions MovieWeekdaySessions) bool {
	for _, sessions := range movieSessions {
		if len(sessions) == 0 {
			return false
		}

		for _, session := range sessions {
			if (session.Hour < 13 || session.Hour > 22.5) || len(session.Room) == 0 {
				return false
			}
		}
	}

	return true
}

func validateRate(rate string) bool {
	for _, r := range Rates {
		if r == rate {
			return true
		}
	}
	return false
}
