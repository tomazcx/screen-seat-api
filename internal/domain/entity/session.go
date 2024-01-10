package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID          string    `json:"ID"`
	Movie       Movie     `json:"Movie"`
	DateTime    time.Time `json:"DateTime"`
	Room        string    `json:"Room"`
	IsSubtitled bool      `json:"IsSubtitled"`
}

func NewSession(movie Movie, dateTime time.Time, room string, isSubtitled bool) *Session {
	return &Session{
		ID:          uuid.New().String(),
		Movie:       movie,
		DateTime:    dateTime,
		Room:        room,
		IsSubtitled: isSubtitled,
	}
}
