package service

import (
	"context"
)

func (s *service[T]) GetAll(ctx context.Context) ([]T, error) {
	return s.repo.GetAll(ctx)
}

func (s *service[T]) GetByID(ctx context.Context, id string) (T, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service[T]) Create(ctx context.Context, entity T) error {
	return s.repo.Create(ctx, entity)
}

func (s *service[T]) Update(ctx context.Context, id string, entity T) error {
	return s.repo.Update(ctx, id, entity)
}

func (s *service[T]) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *service[T]) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error) {
	return s.repo.GetByFilter(ctx, filter)
}
