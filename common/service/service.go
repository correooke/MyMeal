package service

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
	repo "github.com/correooke/MyMeal/common/repository"
)

type service[T model.Entity] struct {
	repo repo.CommonRepository[T]
}

func NewService[T model.Entity](repo repo.CommonRepository[T]) CommonService[T] {
	return &service[T]{repo: repo}
}

type CommonService[T model.Entity] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, id string, entity T) error
	Delete(ctx context.Context, id string) error
}
