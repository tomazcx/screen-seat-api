package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (m *CategoryRepositoryMock) Exists(category string) (bool, error){
	args :=  m.Called(category)
	return args.Bool(0), args.Error(1)
}

func (m *CategoryRepositoryMock) FindAll() ([]entity.Category, error){
	args := m.Called()
	return args.Get(0).([]entity.Category), args.Error(1)
}

func (m *CategoryRepositoryMock) Create(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *CategoryRepositoryMock) Delete(id string) error{ 
	args := m.Called(id)
	return args.Error(0)
}

