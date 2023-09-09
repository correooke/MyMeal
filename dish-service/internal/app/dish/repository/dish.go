package repository

import (
	"context"
	"dish-service/api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dishCollection = "dishes"
)

type DishRepository interface {
	GetDishes(ctx context.Context) ([]model.Dish, error)
	GetDish(ctx context.Context, id string) (model.Dish, error)
	CreateDish(ctx context.Context, dish model.Dish) error
	UpdateDish(ctx context.Context, id string, dish model.Dish) error
	DeleteDish(ctx context.Context, id string) error
}

type mongoDishRepository struct {
	client   *mongo.Client
	database string
}

func NewMongoDishRepository(client *mongo.Client, database string) DishRepository {
	return &mongoDishRepository{client: client, database: database}
}

func (m *mongoDishRepository) collection() *mongo.Collection {
	return m.client.Database(m.database).Collection(dishCollection)
}

func (m *mongoDishRepository) GetDishes(ctx context.Context) ([]model.Dish, error) {
	cur, err := m.collection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var dishes []model.Dish
	if err = cur.All(ctx, &dishes); err != nil {
		return nil, err
	}
	return dishes, nil
}

func (m *mongoDishRepository) GetDish(ctx context.Context, id string) (model.Dish, error) {
	var dish model.Dish
	if err := m.collection().FindOne(ctx, bson.M{"_id": id}).Decode(&dish); err != nil {
		return model.Dish{}, err
	}
	return dish, nil
}

func (m *mongoDishRepository) CreateDish(ctx context.Context, dish model.Dish) error {
	_, err := m.collection().InsertOne(ctx, dish)
	return err
}

func (m *mongoDishRepository) UpdateDish(ctx context.Context, id string, dish model.Dish) error {
	_, err := m.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": dish})
	return err
}

func (m *mongoDishRepository) DeleteDish(ctx context.Context, id string) error {
	_, err := m.collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}
