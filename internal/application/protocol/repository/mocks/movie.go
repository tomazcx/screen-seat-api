package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)


type MovieRepositoryMock struct {
	mock.Mock
}

func (m *MovieRepositoryMock) Exists(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MovieRepositoryMock) FindById(id string) (*entity.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Movie), args.Error(1)
}

func (m *MovieRepositoryMock) FindAllExhbition() ([]entity.Movie, error) {
	args := m.Called()
	return args.Get(0).([]entity.Movie), args.Error(1)
}

func (m *MovieRepositoryMock) FindMany(page int, limit int, sort string, title string, rate string, category string) ([]entity.Movie, error) {
	args := m.Called(page, limit, sort, title, rate, category)
	return args.Get(0).([]entity.Movie), args.Error(1)
}

func (m *MovieRepositoryMock) Create(movie *entity.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func (m *MovieRepositoryMock) Update(movie *entity.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func (m *MovieRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(1)
}

