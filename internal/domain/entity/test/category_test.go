package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestNewCategory(t *testing.T) {
	category := entity.NewCategory("Action")
	
	assert.NotNil(t, category.ID)
	assert.Equal(t, category.Name, "Action")
}
