package uow

import "context"

type Command interface {
	Execute(ctx context.Context) error
}
