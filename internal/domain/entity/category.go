package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(category string) *Category {
	return &Category{
		ID: uuid.New().String(),
		Name: category,
	}
}
