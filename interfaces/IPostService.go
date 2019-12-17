package interfaces

import "github.com/djamboe/mtools-post-service/models"

type IPostService interface {
	CreatePostProcess(postParamModel models.PostModel) (interface{}, error)
	UpdatePostProcess(id string, postParamModel models.PostModel) (interface{}, error)
}
