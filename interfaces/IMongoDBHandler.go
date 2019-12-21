package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoDBHandler interface {
	Find(filter bson.M, collectionName string, databaseName string) (*mongo.Cursor, error)
	FindOne(filter bson.M, collectionName string, databaseName string) (IRowMongoDB, error)
	InsertOne(data interface{}, collectionName string, databaseName string) (interface{}, error)
	UpdateOne(id string, data interface{}, collectionName string, databaseName string) (interface{}, error)
	UpdateMany(data interface{}, collectionName string, databaseName string, filterParam bson.M) (interface{}, error)
}
type IRowMongoDB interface {
	DecodeResults(v interface{}) error
}

type ICursorMongoDB interface {
	Next() *mongo.Cursor
	DecodeResults(v interface{}) error
}
