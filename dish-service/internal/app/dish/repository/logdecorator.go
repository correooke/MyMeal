package repository

import (
	"context"
	"dish-service/api/model"

	"github.com/sirupsen/logrus"
)

type LogDishRepositoryDecorator struct {
	repo DishRepository
}

func NewLogDishRepositoryDecorator(repo DishRepository) DishRepository {
	return &LogDishRepositoryDecorator{repo: repo}
}

func (d *LogDishRepositoryDecorator) GetDishes(ctx context.Context) ([]model.Dish, error) {
	logrus.Info("GetDishes called")
	dishes, err := d.repo.GetDishes(ctx)
	if err != nil {
		logrus.Errorf("Error in GetDishes: %v", err)
	}
	return dishes, err
}

func (d *LogDishRepositoryDecorator) GetDish(ctx context.Context, id string) (model.Dish, error) {
	logrus.Infof("GetDish called with ID: %s", id)
	dish, err := d.repo.GetDish(ctx, id)
	if err != nil {
		logrus.Errorf("Error in GetDish: %v", err)
	}
	return dish, err
}

func (d *LogDishRepositoryDecorator) CreateDish(ctx context.Context, dish model.Dish) error {
	logrus.Infof("CreateDish called with Name: %s", dish.Name)
	err := d.repo.CreateDish(ctx, dish)
	if err != nil {
		logrus.Errorf("Error in CreateDish: %v", err)
	}
	return err
}

func (d *LogDishRepositoryDecorator) UpdateDish(ctx context.Context, id string, dish model.Dish) error {
	logrus.Infof("UpdateDish called with ID: %s", id)
	err := d.repo.UpdateDish(ctx, id, dish)
	if err != nil {
		logrus.Errorf("Error in UpdateDish: %v", err)
	}
	return err
}

func (d *LogDishRepositoryDecorator) DeleteDish(ctx context.Context, id string) error {
	logrus.Infof("DeleteDish called with ID: %s", id)
	err := d.repo.DeleteDish(ctx, id)
	if err != nil {
		logrus.Errorf("Error in DeleteDish: %v", err)
	}
	return err
}
