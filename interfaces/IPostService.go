package interfaces

import "github.com/djamboe/mtools-post-service/models"

type IPostService interface {
	CreatePostProcess(postParamModel models.PostModel) (interface{}, error)
	UpdatePostProcess(id string, postParamModel models.PostModel) (interface{}, error)
	CreatePostDetailProcess(postParamModel models.PostDetailModel) (interface{}, error)
	UpdatePostDetailProcess(id string, postParamModel models.PostDetailModel) (interface{}, error)
	GetPostDataProcess(postParamModel models.PostDataParamModel) (models.PostModel, error)
	GetPostDetailDataProcess(postParamModel models.GetPostDetailParamModel) (models.PostDetailModel, error)
	GetListPostDataProcess(postParamModel models.GetListPostDataParam) ([]*models.PostModel, error)
	GetListPostDataDetailProcess(postParamModel models.GetListPostDataDetailParam) ([]*models.PostDetailModel, error)
	DeletePostDataProcess(id string, postParamModel models.DeletePostModel) (interface{}, error)
	DeletePostDataDetailProcess(id string, postParamModel models.DeletePostModel) (interface{}, error)
	GetListWeeklyPlanDataProcess(dataWeeklyPlanParamModel models.GetWeeklyPlanParamModel) ([]*models.WeeklyPlan, error)
}
