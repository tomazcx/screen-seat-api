-- name: CategoryExists :one
SELECT COUNT(1) FROM category WHERE id=$1;

-- name: FindManyCategories :many
SELECT * FROM category;

-- name: CreateCategory :exec
INSERT INTO category (id, name) VALUES ($1, $2);

-- name: DeleteCategory :exec
DELETE FROM category WHERE id=$1;
