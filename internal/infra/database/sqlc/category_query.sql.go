// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: category_query.sql

package query

import (
	"context"

	"github.com/google/uuid"
)

const categoryExists = `-- name: CategoryExists :one
SELECT COUNT(1) FROM category WHERE id=$1
`

func (q *Queries) CategoryExists(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, categoryExists, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createCategory = `-- name: CreateCategory :exec
INSERT INTO category (id, name) VALUES ($1, $2)
`

type CreateCategoryParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category WHERE id=$1
`

func (q *Queries) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const findManyCategories = `-- name: FindManyCategories :many
SELECT id, name, created_at FROM category
`

func (q *Queries) FindManyCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, findManyCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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