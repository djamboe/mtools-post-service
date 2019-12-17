package models

import (
	"time"
)

type PostModelParam struct {
	DbId         string    `json:"_id,omitempty"`
	CustomerId   int64     `json:"customerId"`
	CustomerName string    `json:"customerName"`
	UserId       int64     `json:"userId"`
	Chanel       string    `json:"chanel"`
	Description  string    `json:"description"`
	Product      string    `json:"product"`
	Phone        string    `json:"phone"`
	Pic          string    `json:"pic"`
	Price        float64   `json:"price"`
	Notes        string    `json:"notes"`
	Status       int32     `json:"status"`
	Photo        []Photo   `json:"photo"`
	CreatedOn    time.Time `json:"createdOn"`
	UpdatedOn    time.Time `json:"updatedOn"`
}
