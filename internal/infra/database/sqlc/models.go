// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package query

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

type JoinMovieCategory struct {
	MovieID    uuid.UUID
	CategoryID uuid.UUID
}

type Movie struct {
	ID          uuid.UUID
	Title       string
	Description sql.NullString
	Duration    int32
	Poster      sql.NullString
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
}
