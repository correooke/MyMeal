package service

import (
	"context"
	"dish-service/api/model"
	"dish-service/internal/app/dish/repository"
)

type dishService struct {
	repo repository.DishRepository
}

func NewDishService(repo repository.DishRepository) DishService {
	return &dishService{repo: repo}
}

type DishService interface {
	GetDishes(ctx context.Context) ([]model.Dish, error)
	GetDish(ctx context.Context, id string) (model.Dish, error)
	CreateDish(ctx context.Context, dish model.Dish) error
	UpdateDish(ctx context.Context, id string, dish model.Dish) error
	DeleteDish(ctx context.Context, id string) error
}
