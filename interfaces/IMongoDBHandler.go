package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
)

type IMongoDBHandler interface {
	FindOne(filter bson.M, collectionName string, databaseName string) (IRowMongoDB, error)
	InsertOne(data interface{}, collectionName string, databaseName string) (interface{}, error)
}
type IRowMongoDB interface {
	Next() bool
	DecodeResults(v interface{}) error
}
