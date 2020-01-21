package models

import "time"

type PostModel struct {
	CustomerId   string    `json:"customerId"`
	ProductId    string    `json:"productId"`
	ProductName  string    `json:"productName"`
	CustomerName string    `json:"customerName"`
	UserId       string    `json:"userId"`
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
	IsDeleted    bool      `json:"isDeleted"`
}
