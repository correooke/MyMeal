package service

import (
	"context"
	"dish-service/api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRepository es un mock del repositorio de Dish para los tests.
type MockRepository struct{}

// CreateDish implements repository.DishRepository.
func (*MockRepository) CreateDish(ctx context.Context, dish model.Dish) error {
	panic("unimplemented")
}

// DeleteDish implements repository.DishRepository.
func (*MockRepository) DeleteDish(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetDish implements repository.DishRepository.
func (*MockRepository) GetDish(ctx context.Context, id string) (model.Dish, error) {
	panic("unimplemented")
}

// UpdateDish implements repository.DishRepository.
func (*MockRepository) UpdateDish(ctx context.Context, id string, dish model.Dish) error {
	panic("unimplemented")
}

func (m *MockRepository) GetDishes(ctx context.Context) ([]model.Dish, error) {
	return []model.Dish{{Name: "TestDish", Description: "TestDescription", Price: 10.99}}, nil
}

//... (puedes agregar mocks para los otros métodos también) ...

func TestGetDishes(t *testing.T) {
	repo := &MockRepository{}
	service := NewDishService(repo)

	dishes, err := service.GetDishes(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dishes))
	assert.Equal(t, "TestDish", dishes[0].Name)
}

//... (puedes agregar tests para los otros métodos también) ...
