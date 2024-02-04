package uow

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
)

type UpdateFunc[T model.Entity] func(ctx context.Context, id string, entity T) error

type UpdateCommand[T model.Entity] struct {
	UpdateFunc UpdateFunc[T]
	Entity     T
}

func (cmd *UpdateCommand[T]) Execute(ctx context.Context) error {
	return cmd.UpdateFunc(ctx, cmd.Entity.GetID(), cmd.Entity)
}
