package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeeklyPlan struct {
	DbId   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Date   primitive.DateTime `json:"date"`
	Title  string             `json:"title"`
	Status string             `json:"status"`
}
