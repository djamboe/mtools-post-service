package models

type PostModel struct {
	Id           int64  `json:"id"`
	DbId         string `json:"_id"`
	CustomerId   int64  `json:"customerId"`
	CustomerName string `json:"customerName"`
	UserId       int64  `json:"userId"`
	Chanel       string `json:"chanel"`
	Description  string `json:"description"`
	Product      string `json:"product"`
	Phone        string `json:"phone"`
	Pic          string `json:"pic"`
	Price        int64  `json:"price"`
	Notes        string `json:"notes"`
	Status       int32  `json:"status"`
	CreatedOn    int64  `json:"createdOn"`
	UpdatedOn    int64  `json:"updatedOn"`
	Photo        Photo  `json:"photo"`
}
