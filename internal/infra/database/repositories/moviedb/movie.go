package moviedb

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
	query "github.com/tomazcx/screen-seat-api/internal/infra/database/sqlc"
)

type MovieRepository struct {
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
		ID: row.ID.String(),
		Title: row.Title,
		Description: row.Description.String,
		Poster: row.Poster.String,
		AgeRating: row.AgeRating,
		Duration: int(row.Duration),
		StartDate: row.StartDate,
		EndDate: row.EndDate,
		Categories: categoriesSlice,
	}

	return movie, nil	
}

func (r *MovieRepository) FindAllInExhibition() ([]entity.Movie, error){


}
