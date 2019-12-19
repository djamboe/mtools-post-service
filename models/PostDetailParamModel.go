package models

import "time"

type PostDetailParamModel struct {
	DbId        string    `json:"_id"`
	PostId      string    `json:"postId"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Status      int32     `json:"status"`
	CreatedOn   time.Time `json:"createdOn"`
	UpdatedOn   time.Time `json:"updatedOn"`
	Photo       []Photo   `json:"photo"`
	IsDeleted   bool      `json:"isDeleted"`
}
