package repository

import (
	"context"

	"github.com/correooke/MyMeal/common/db"
	"github.com/correooke/MyMeal/common/model"
	"github.com/correooke/MyMeal/common/repository/uow"
)

type mongoRepositoryUoW[T model.Entity] struct {
	mongoRepository[T]
	uoW uow.UnitOfWork
}

func NewMongoRepositoryUoW[T model.Entity](collection db.CollectionInterface, uoW uow.UnitOfWork) CommonRepository[T] {
	r := NewMongoRepository[T](collection)
	return &mongoRepositoryUoW[T]{mongoRepository: r.(mongoRepository[T]), uoW: uoW}
}

func (m mongoRepositoryUoW[T]) Create(ctx context.Context, entity T) error {
	cmd := &uow.InsertCommand[T]{CreateFunc: m.mongoRepository.Create, Document: entity}
	m.uoW.AddCommand(cmd)
	return nil
}

func (m mongoRepositoryUoW[T]) Update(ctx context.Context, id string, entity T) error {
	cmd := &uow.UpdateCommand[T]{UpdateFunc: m.mongoRepository.Update, Entity: entity}
	m.uoW.AddCommand(cmd)
	return nil
}

func (m mongoRepositoryUoW[T]) Delete(ctx context.Context, id string) error {
	cmd := &uow.DeleteCommand{DeleteFunc: m.mongoRepository.Delete, ID: id}
	m.uoW.AddCommand(cmd)
	return nil
}
