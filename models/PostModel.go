package models

type PostModel struct {
	Id           int64  `json:"id"`
	DbId         string `json:"_id"`
	CustomerId   int64  `json:"customerId"`
	CustomerName string `json:"customerName"`
	UserId       int64  `json:"userId"`
	Chanel       string `json:"chanel"`
	Description  string `json:"description"`
}
