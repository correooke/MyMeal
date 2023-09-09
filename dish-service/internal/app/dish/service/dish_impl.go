package service

import (
	"context"
	"dish-service/api/model"
)

func (s *dishService) GetDishes(ctx context.Context) ([]model.Dish, error) {
	return s.repo.GetDishes(ctx)
}

func (s *dishService) GetDish(ctx context.Context, id string) (model.Dish, error) {
	return s.repo.GetDish(ctx, id)
}

func (s *dishService) CreateDish(ctx context.Context, dish model.Dish) error {
	return s.repo.CreateDish(ctx, dish)
}

func (s *dishService) UpdateDish(ctx context.Context, id string, dish model.Dish) error {
	return s.repo.UpdateDish(ctx, id, dish)
}

func (s *dishService) DeleteDish(ctx context.Context, id string) error {
	return s.repo.DeleteDish(ctx, id)
}
