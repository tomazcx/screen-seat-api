package entity

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(category string) *Category {
	return &Category{
		ID: uuid.New().String(),
		Name: category,
	}
}
