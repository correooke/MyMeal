package repository

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
)

type CommonRepository[T model.Entity] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, id string, entity T) error
	Delete(ctx context.Context, id string) error
}
