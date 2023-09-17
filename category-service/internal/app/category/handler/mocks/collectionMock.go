package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionMock struct {
	mock.Mock
}

func (m *CollectionMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *CollectionMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *CollectionMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *CollectionMock) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(filter)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *CollectionMock) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}
