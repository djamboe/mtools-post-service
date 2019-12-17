package interfaces

import (
	"github.com/djamboe/mtools-post-service/models"
)

type IPostRepository interface {
	CreatePost(postParamModels models.PostModel) (interface{}, error)
	UpdatePost(id string, updateParamModels models.PostModel) (interface{}, error)
}
