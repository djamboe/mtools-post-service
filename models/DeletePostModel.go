package models

import "time"

type DeletePostModel struct {
	IsDeleted bool      `json:"isdeleted"`
	UpdatedOn time.Time `json:"updated_on"`
}
