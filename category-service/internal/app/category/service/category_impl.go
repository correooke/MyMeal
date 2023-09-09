package service

import (
	"category-service/api/model"
	"context"
)

func (s *categoryService) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *categoryService) GetCategory(ctx context.Context, id string) (model.Category, error) {
	return s.repo.GetCategory(ctx, id)
}

func (s *categoryService) CreateCategory(ctx context.Context, category model.Category) error {
	return s.repo.CreateCategory(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id string, category model.Category) error {
	return s.repo.UpdateCategory(ctx, id, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id string) error {
	return s.repo.DeleteCategory(ctx, id)
}
