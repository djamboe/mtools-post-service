package infrastructures

import (
	"context"
	"fmt"
	"github.com/djamboe/mtools-post-service/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandler struct {
	Conn *mongo.Client
}

func (handler *MongoDBHandler) FindOne(filter bson.M, collectionName string, dbName string) (interfaces.IRowMongoDB, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows := collection.FindOne(context.TODO(), filter)
	row := new(MongoRow)
	row.Rows = rows
	return row, nil
}

func (handler *MongoDBHandler) InsertOne(data interface{}, collectionName string, dbName string) (interface{}, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
	}
	return insertResult, nil
}

func (handler *MongoDBHandler) UpdateOne(id string, data interface{}, collectionName string, dbName string) (interface{}, error) {
	docId := id
	objId, err := primitive.ObjectIDFromHex(docId)

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.D{{Key: "$set", Value: data}}

	collection := handler.Conn.Database(dbName).Collection(collectionName)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err)
	}
	return updateResult, nil
}

func (handler *MongoDBHandler) Find(filter bson.M, collectionName string, dbName string) (*mongo.Cursor, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows, _ := collection.Find(context.TODO(), filter)
	row := new(MongoCursor)
	row.Rows = rows
	return rows, nil
}

type MongoCursor struct {
	Rows *mongo.Cursor
}

func (r MongoCursor) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}

func (r MongoCursor) Next() *mongo.Cursor {
	return r.Next()
}

type MongoRow struct {
	Rows *mongo.SingleResult
}

func (r *MongoRow) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}
