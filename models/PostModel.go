package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostModel struct {
	DbId         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerId   string             `json:"customerId"`
	ProductId    string             `json:"productId"`
	CustomerName string             `json:"customerName"`
	UserId       string             `json:"userId"`
	Chanel       string             `json:"chanel"`
	Description  string             `json:"description"`
	Product      string             `json:"product"`
	Phone        string             `json:"phone"`
	Pic          string             `json:"pic"`
	Price        float64            `json:"price"`
	Notes        string             `json:"notes"`
	Status       string             `json:"status"`
	CreatedOn    time.Time          `json:"createdOn"`
	UpdatedOn    time.Time          `json:"updatedOn"`
	IsDeleted    bool               `json:"isDeleted"`
}
