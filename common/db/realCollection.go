package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RealCollection es una implementación real de CollectionInterface usando mongo.Collection.
type RealCollection struct {
	collection *mongo.Collection
}

// NewRealCollection crea una nueva instancia de RealCollection.
func NewRealCollection(collection *mongo.Collection) *RealCollection {
	return &RealCollection{collection: collection}
}

// InsertOne inserta un documento en la colección.
func (rc *RealCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return rc.collection.InsertOne(ctx, document, opts...)
}

// UpdateOne actualiza un documento en la colección que coincida con el filtro.
func (rc *RealCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return rc.collection.UpdateOne(ctx, filter, update, opts...)
}

// DeleteOne elimina un documento de la colección que coincida con el filtro.
func (rc *RealCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return rc.collection.DeleteOne(ctx, filter, opts...)
}

// Find busca documentos en la colección que coincidan con el filtro.
func (rc *RealCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return rc.collection.Find(ctx, filter, opts...)
}

// FindOne busca un documento en la colección que coincida con el filtro.
func (rc *RealCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return rc.collection.FindOne(ctx, filter, opts...)
}
