package uow

import (
	"context"
)

type DeleteFunc func(ctx context.Context, id string) error

type DeleteCommand struct {
	DeleteFunc DeleteFunc
	ID         string
}

func (cmd *DeleteCommand) Execute(ctx context.Context) error {
	return cmd.DeleteFunc(ctx, cmd.ID)
}
