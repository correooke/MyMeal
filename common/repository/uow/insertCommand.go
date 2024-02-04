package uow

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
)

type CreateFunc[T model.Entity] func(ctx context.Context, entity T) error

type InsertCommand[T model.Entity] struct {
	CreateFunc CreateFunc[T]
	Document   T
}

func (cmd *InsertCommand[T]) Execute(ctx context.Context) error {
	return cmd.CreateFunc(ctx, cmd.Document)
}
