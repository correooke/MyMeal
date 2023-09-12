package service

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
)

func (s *service) GetAll(ctx context.Context) ([]model.Entity, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id string) (model.Entity, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, entity model.Entity) error {
	return s.repo.Create(ctx, entity)
}

func (s *service) Update(ctx context.Context, id string, entity model.Entity) error {
	return s.repo.Update(ctx, id, entity)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]model.Entity, error) {
	return s.repo.GetByFilter(ctx, filter)
}
