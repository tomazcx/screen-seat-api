-- name: MovieExists :one
SELECT COUNT(1) FROM movie WHERE id = $1;

-- name: JoinMovieCategory :exec
INSERT INTO join_movie_category (movie_id, category_name) VALUES ($1, $2);

-- name: FindMovieById :one
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
    m.id;

-- name: FindExhibitionMovies :many

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
WHERE startDate < CURRENT_DATE
AND endDate > CURRENT_DATE
GROUP BY
    m.id;

-- name: FindManyMovies :many
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
OFFSET $2;

-- name: CreateMovie :exec
INSERT INTO movie (id, title, duration, description, age_rating, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateMovie :exec
UPDATE movie SET title=$2,duration=$3, description=$4, age_rating=$5, start_date=$6, end_date=$7 WHERE id = $1;

-- name: UpdateMoviePoster :exec
UPDATE movie SET poster=$2 WHERE id=$1;

-- name: DeleteMovie :exec
DELETE FROM movie WHERE id = $1;

