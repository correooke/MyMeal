package repository

import (
	"context"

	"github.com/correooke/MyMeal/common/db"
	"github.com/correooke/MyMeal/common/model"
	"go.mongodb.org/mongo-driver/bson"
)

type CommonRepository[T model.Entity] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, id string, entity T) error
	Delete(ctx context.Context, id string) error
}

type mongoRepository[T model.Entity] struct {
	collection db.CollectionInterface
}

func NewMongoRepository[T model.Entity](collection db.CollectionInterface) CommonRepository[T] {
	return &mongoRepository[T]{collection: collection}
}

func (m *mongoRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var collection []T
	if err = cur.All(ctx, &collection); err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *mongoRepository[T]) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error) {
	var entities []T

	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entity T
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}

func (m *mongoRepository[T]) GetByID(ctx context.Context, id string) (T, error) {
	var entity T
	if err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity); err != nil {
		return entity, err
	}
	return entity, nil
}

func (m *mongoRepository[T]) Create(ctx context.Context, entity T) error {
	_, err := m.collection.InsertOne(ctx, entity)
	return err
}

func (m *mongoRepository[T]) Update(ctx context.Context, id string, entity T) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": entity})
	return err
}

func (m *mongoRepository[T]) Delete(ctx context.Context, id string) error {
	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
