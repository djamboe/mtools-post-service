package interfaces

import (
	"github.com/djamboe/mtools-post-service/models"
)

type IPostRepository interface {
	CreatePost(postParamModels models.PostModel) (interface{}, error)
	UpdatePost(id string, updateParamModels models.PostModel) (interface{}, error)
	CreatePostDetail(postParamModels models.PostDetailModel) (interface{}, error)
	UpdatePostDetail(id string, updateParamModels models.PostDetailModel) (interface{}, error)
	GetPostDataById(dataPostParamModels models.PostDataParamModel) (models.PostModel, error)
}
