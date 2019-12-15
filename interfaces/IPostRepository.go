package interfaces

import (
	"github.com/djamboe/mtools-post-service/models"
)

type IPostRepository interface {
	CreatePost(postParamModels models.PostModelParam) (interface{}, error)
}
