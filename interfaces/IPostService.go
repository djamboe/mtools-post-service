package interfaces

import "github.com/djamboe/mtools-post-service/models"

type IPostService interface {
	CreatePostProcess(postParamModel models.PostModelParam) (interface{}, error)
}
