package service

import (
	"category-service/api/model"
	"category-service/internal/app/category/repository"
	"context"
)

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

type CategoryService interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, id string) (model.Category, error)
	CreateCategory(ctx context.Context, category model.Category) error
	UpdateCategory(ctx context.Context, id string, category model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}
