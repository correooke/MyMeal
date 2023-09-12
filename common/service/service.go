package service

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
	repo "github.com/correooke/MyMeal/common/repository"
)

type service struct {
	repo repo.CommonRepository
}

func NewService(repo repo.CommonRepository) CommonService {
	return &service{repo: repo}
}

type CommonService interface {
	GetAll(ctx context.Context) ([]model.Entity, error)
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]model.Entity, error)
	GetByID(ctx context.Context, id string) (model.Entity, error)
	Create(ctx context.Context, entity model.Entity) error
	Update(ctx context.Context, id string, entity model.Entity) error
	Delete(ctx context.Context, id string) error
}
