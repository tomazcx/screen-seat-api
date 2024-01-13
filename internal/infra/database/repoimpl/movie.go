package repoimpl

import (
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
	query "github.com/tomazcx/screen-seat-api/internal/infra/database/sqlc"
)

type MovieRepository struct {
	db      *sql.DB
	queries query.Queries
}

func (r *MovieRepository) Exists(id string) (bool, error) {
	count, err := r.queries.MovieExists(context.Background(), uuid.MustParse(id))

	if err != nil {
		return false, err
	}

	return count > 0, err
}

func (r *MovieRepository) FindById(id string) (*entity.Movie, error) {
	row, err := r.queries.FindMovieById(context.Background(), uuid.MustParse(id))

	if err != nil {
		return nil, err
	}

	categoriesSlice := strings.Split(string(row.Categories), ",")

	movie := &entity.Movie{
		ID:          row.ID.String(),
		Title:       row.Title,
		Description: row.Description.String,
		Poster:      row.Poster.String,
		AgeRating:   row.AgeRating,
		Duration:    int(row.Duration),
		StartDate:   row.StartDate,
		EndDate:     row.EndDate,
		Categories:  categoriesSlice,
	}

	return movie, nil
}

func (r *MovieRepository) FindAllExhibition() ([]entity.Movie, error) {
	rows, err := r.queries.FindExhibitionMovies(context.Background())

	if err != nil {
		return nil, err
	}

	var movies []entity.Movie

	for _, row := range rows {
		categoriesSlice := strings.Split(string(row.Categories), ",")
		movie := entity.Movie{
			ID:          row.ID.String(),
			Title:       row.Title,
			Description: row.Description.String,
			Poster:      row.Poster.String,
			AgeRating:   row.AgeRating,
			Duration:    int(row.Duration),
			StartDate:   row.StartDate,
			EndDate:     row.EndDate,
			Categories:  categoriesSlice,
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) FindMany(page int, limit int, sort string, title string, rate string, category string) ([]entity.Movie, error) {

	params := query.FindManyMoviesParams{
		Limit:   int32(limit),
		Offset:  int32((page - 1) * limit),
		Column3: sort,
		Title:   "%" + title + "%",
		Column5: rate,
		Column6: category,
	}
	rows, err := r.queries.FindManyMovies(context.Background(), params)
	if err != nil {
		return nil, err
	}

	var movies []entity.Movie

	for _, row := range rows {
		categoriesSlice := strings.Split(string(row.Categories), ",")
		movie := entity.Movie{
			ID:          row.ID.String(),
			Title:       row.Title,
			Description: row.Description.String,
			Poster:      row.Poster.String,
			AgeRating:   row.AgeRating,
			Duration:    int(row.Duration),
			StartDate:   row.StartDate,
			EndDate:     row.EndDate,
			Categories:  categoriesSlice,
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) Create(movie *entity.Movie) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	queryTx := r.queries.WithTx(tx)
	params := query.CreateMovieParams{
		ID:          uuid.MustParse(movie.ID),
		Title:       movie.Title,
		Description: sql.NullString{String: movie.Description},
		AgeRating:   movie.AgeRating,
		Duration:    int32(movie.Duration),
		StartDate:   movie.StartDate,
		EndDate:     movie.EndDate,
	}
	err = queryTx.CreateMovie(context.Background(), params)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, category := range movie.Categories {
		params := query.JoinMovieCategoryParams{
			MovieID:      uuid.MustParse(movie.ID),
			CategoryName: category,
		}
		err = queryTx.JoinMovieCategory(context.Background(), params)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *MovieRepository) UpdatePoster(id string, posterUrl string) error {
	params := query.UpdateMoviePosterParams{
		ID:     uuid.MustParse(id),
		Poster: sql.NullString{String: posterUrl},
	}
	return r.queries.UpdateMoviePoster(context.Background(), params)
}

func (r *MovieRepository) Update(movie *entity.Movie) error {
	params := query.UpdateMovieParams{
		ID:          uuid.MustParse(movie.ID),
		Title:       movie.Title,
		Description: sql.NullString{String: movie.Description},
		AgeRating:   movie.AgeRating,
		Duration:    int32(movie.Duration),
		StartDate:   movie.StartDate,
		EndDate:     movie.EndDate,
	}
	return r.queries.UpdateMovie(context.Background(), params)
}

func (r *MovieRepository) Delete(id string) error {
	return r.queries.DeleteMovie(context.Background(), uuid.MustParse(id))
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	queries := query.New(db)
	return &MovieRepository{
		queries: *queries,
	}
}
