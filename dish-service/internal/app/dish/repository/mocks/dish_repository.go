package mocks

import (
	"context"
	"dish-service/api/model"

	"github.com/stretchr/testify/mock"
)

type DishRepositoryMock struct {
	mock.Mock
}

func (m *DishRepositoryMock) GetDishes(ctx context.Context) ([]model.Dish, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Dish), args.Error(1)
}

//... (implementa los otros métodos mockeados de la misma manera) ...
