package repository

import (
	"context"

	"github.com/correooke/MyMeal/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommonRepository interface {
	GetAll(ctx context.Context) ([]model.Entity, error)
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]model.Entity, error)
	GetByID(ctx context.Context, id string) (model.Entity, error)
	Create(ctx context.Context, entity model.Entity) error
	Update(ctx context.Context, id string, entity model.Entity) error
	Delete(ctx context.Context, id string) error
}

type mongoRepository struct {
	client         *mongo.Client
	database       string
	collectionName string
}

func NewMongoRepository(client *mongo.Client, database string, collectionName string) CommonRepository {
	return &mongoRepository{client: client, database: database, collectionName: collectionName}
}

func (m *mongoRepository) collection() *mongo.Collection {
	return m.client.Database(m.database).Collection(m.collectionName)
}

func (m *mongoRepository) GetAll(ctx context.Context) ([]model.Entity, error) {
	cur, err := m.collection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var collection []model.Entity
	if err = cur.All(ctx, &collection); err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *mongoRepository) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]model.Entity, error) {
	var entities []model.Entity

	cursor, err := m.collection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entity model.Entity
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

func (m *mongoRepository) GetByID(ctx context.Context, id string) (model.Entity, error) {
	var entity model.Entity
	if err := m.collection().FindOne(ctx, bson.M{"_id": id}).Decode(&entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (m *mongoRepository) Create(ctx context.Context, entity model.Entity) error {
	_, err := m.collection().InsertOne(ctx, entity)
	return err
}

func (m *mongoRepository) Update(ctx context.Context, id string, entity model.Entity) error {
	_, err := m.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": entity})
	return err
}

func (m *mongoRepository) Delete(ctx context.Context, id string) error {
	_, err := m.collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}
