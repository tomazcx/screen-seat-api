// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: movie_query.sql

package query

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createMovie = `-- name: CreateMovie :exec
INSERT INTO movie (id, title, duration, description, age_rating, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Duration    int32
	Description sql.NullString
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) error {
	_, err := q.db.ExecContext(ctx, createMovie,
		arg.ID,
		arg.Title,
		arg.Duration,
		arg.Description,
		arg.AgeRating,
		arg.StartDate,
		arg.EndDate,
	)
	return err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movie WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, id)
	return err
}

const findManyMovies = `-- name: FindManyMovies :many
SELECT
    m.id,
    m.title,
    m.description,
    m.duration,
    m.poster,
    m.age_rating,
    m.start_date,
    m.end_date,
    STRING_AGG(c.name, ',') AS categories
FROM
    movie m
INNER JOIN
    join_movie_category j ON m.id = j.movie_id
INNER JOIN
    category c ON j.category_id = c.id
WHERE
    m.title ILIKE $4
AND 
   ($5 = '' OR m.age_rating = $5)
AND
   ($6 = '' OR c.name = $6)
GROUP BY
    m.id
ORDER BY $3
LIMIT $1
OFFSET $2
`

type FindManyMoviesParams struct {
	Limit   int32
	Offset  int32
	Column3 interface{}
	Title   string
	Column5 interface{}
	Column6 interface{}
}

type FindManyMoviesRow struct {
	ID          uuid.UUID
	Title       string
	Description sql.NullString
	Duration    int32
	Poster      sql.NullString
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	Categories  []byte
}

func (q *Queries) FindManyMovies(ctx context.Context, arg FindManyMoviesParams) ([]FindManyMoviesRow, error) {
	rows, err := q.db.QueryContext(ctx, findManyMovies,
		arg.Limit,
		arg.Offset,
		arg.Column3,
		arg.Title,
		arg.Column5,
		arg.Column6,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindManyMoviesRow
	for rows.Next() {
		var i FindManyMoviesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Duration,
			&i.Poster,
			&i.AgeRating,
			&i.StartDate,
			&i.EndDate,
			&i.Categories,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findMovieById = `-- name: FindMovieById :one
SELECT
    m.id,
    m.title,
    m.description,
    m.duration,
    m.poster,
    m.age_rating,
    m.start_date,
    m.end_date,
    STRING_AGG(c.name, ',') AS categories
FROM
    movie m
INNER JOIN
    join_movie_category j ON m.id = j.movie_id
INNER JOIN
    category c ON j.category_id = c.id
WHERE
    m.id = $1
GROUP BY
    m.id
`

type FindMovieByIdRow struct {
	ID          uuid.UUID
	Title       string
	Description sql.NullString
	Duration    int32
	Poster      sql.NullString
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
	Categories  []byte
}

func (q *Queries) FindMovieById(ctx context.Context, id uuid.UUID) (FindMovieByIdRow, error) {
	row := q.db.QueryRowContext(ctx, findMovieById, id)
	var i FindMovieByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Duration,
		&i.Poster,
		&i.AgeRating,
		&i.StartDate,
		&i.EndDate,
		&i.Categories,
	)
	return i, err
}

const movieExists = `-- name: MovieExists :one
SELECT COUNT(1) FROM movie WHERE id = $1
`

func (q *Queries) MovieExists(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, movieExists, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateMovie = `-- name: UpdateMovie :exec
UPDATE movie SET title=$2,duration=$3, description=$4, age_rating=$5, start_date=$6, end_date=$7 WHERE id = $1
`

type UpdateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Duration    int32
	Description sql.NullString
	AgeRating   string
	StartDate   time.Time
	EndDate     time.Time
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.ExecContext(ctx, updateMovie,
		arg.ID,
		arg.Title,
		arg.Duration,
		arg.Description,
		arg.AgeRating,
		arg.StartDate,
		arg.EndDate,
	)
	return err
}

const updateMoviePoster = `-- name: UpdateMoviePoster :exec
UPDATE movie SET poster=$2 WHERE id=$1
`

type UpdateMoviePosterParams struct {
	ID     uuid.UUID
	Poster sql.NullString
}

func (q *Queries) UpdateMoviePoster(ctx context.Context, arg UpdateMoviePosterParams) error {
	_, err := q.db.ExecContext(ctx, updateMoviePoster, arg.ID, arg.Poster)
	return err
}
