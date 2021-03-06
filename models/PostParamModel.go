package models

import (
	"time"
)

type PostModelParam struct {
	DbId         string    `json:"_id,omitempty"`
	CustomerId   string    `json:"customerId"`
	CustomerName string    `json:"customerName"`
	UserId       string    `json:"userId"`
	ProductId    string    `json:"productId"`
	Chanel       string    `json:"chanel"`
	Description  string    `json:"description"`
	Product      string    `json:"product"`
	Phone        string    `json:"phone"`
	Pic          string    `json:"pic"`
	Price        float64   `json:"price"`
	Notes        string    `json:"notes"`
	Status       string    `json:"status"`
	CreatedOn    time.Time `json:"createdOn"`
	UpdatedOn    time.Time `json:"updatedOn"`
}
