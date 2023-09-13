package repository

import (
	"category-service/api/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	categoryCollection = "categories"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, id string) (model.Category, error)
	CreateCategory(ctx context.Context, category model.Category) error
	UpdateCategory(ctx context.Context, id string, category model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}

type mongoCategoryRepository struct {
	client   *mongo.Client
	database string
}

func NewMongoCategoryRepository(client *mongo.Client, database string) CategoryRepository {
	return &mongoCategoryRepository{client: client, database: database}
}

func (m *mongoCategoryRepository) collection() *mongo.Collection {
	return m.client.Database(m.database).Collection(categoryCollection)
}

func (m *mongoCategoryRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
	cur, err := m.collection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var categoryes []model.Category
	if err = cur.All(ctx, &categoryes); err != nil {
		return nil, err
	}
	return categoryes, nil
}

func (m *mongoCategoryRepository) GetCategory(ctx context.Context, id string) (model.Category, error) {
	var category model.Category
	if err := m.collection().FindOne(ctx, bson.M{"_id": id}).Decode(&category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (m *mongoCategoryRepository) CreateCategory(ctx context.Context, category model.Category) error {
	_, err := m.collection().InsertOne(ctx, category)
	return err
}

func (m *mongoCategoryRepository) UpdateCategory(ctx context.Context, id string, category model.Category) error {
	_, err := m.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": category})
	return err
}

func (m *mongoCategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	_, err := m.collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}
