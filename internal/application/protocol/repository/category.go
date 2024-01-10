package repository

import "github.com/tomazcx/screen-seat-api/internal/domain/entity"

type ICategoryRepository interface {
	Exists(categoryName string) (bool, error)
	FindAll() ([]entity.Category, error)
	Create(category *entity.Category) error
	Delete(id string) error
}
