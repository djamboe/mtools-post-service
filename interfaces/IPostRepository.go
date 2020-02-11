package interfaces

import (
	"github.com/djamboe/mtools-post-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IPostRepository interface {
	CreatePost(postParamModels models.PostModel) (interface{}, error)
	UpdatePost(id string, updateParamModels models.PostModel) (interface{}, error)
	CreatePostDetail(postParamModels models.PostDetailModel) (interface{}, error)
	UpdatePostDetail(id string, updateParamModels models.PostDetailModel) (interface{}, error)
	GetPostDataById(dataPostParamModels models.PostDataParamModel) (models.PostModel, error)
	GetPostDetailDataById(dataPostParamModels models.GetPostDetailParamModel) (models.PostDetailModel, error)
	GetListPostDataDataByUserId(dataPostParamModels models.GetListPostDataParam) ([]*models.PostModel, error)
	GetListPostDataDetailByPostId(dataPostParamModels models.GetListPostDataDetailParam) ([]*models.PostDetailModel, error)
	DeletePostData(id string, updateParamModels models.DeletePostModel) (interface{}, error)
	DeletePostDataDetail(id string, updateParamModels models.DeletePostModel) (interface{}, error)
	DeleteChildRelationData(collectionName string, updateParamModels models.DeletePostModel, filterParam bson.M) (interface{}, error)
	GetListWeeklyPlanData(dataWeeklyPlanParamModel models.GetWeeklyPlanParamModel) ([]*models.WeeklyPlan, error)
}
