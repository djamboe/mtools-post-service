package services

import (
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

type PostService struct {
	interfaces.IPostRepository
}

func (service *PostService) CreatePostProcess(postParam models.PostModel) (interface{}, error) {
	post, err := service.CreatePost(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) UpdatePostProcess(id string, postParam models.PostModel) (interface{}, error) {
	post, err := service.UpdatePost(id, postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) CreatePostDetailProcess(postParam models.PostDetailModel) (interface{}, error) {
	post, err := service.CreatePostDetail(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) UpdatePostDetailProcess(id string, postParam models.PostDetailModel) (interface{}, error) {
	post, err := service.UpdatePostDetail(id, postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) GetPostDetailData(id string, postParam models.PostDetailModel) (interface{}, error) {
	post, err := service.UpdatePostDetail(id, postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) GetPostDataProcess(postParam models.PostDataParamModel) (models.PostModel, error) {
	post, err := service.GetPostDataById(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) GetPostDetailDataProcess(postParam models.GetPostDetailParamModel) (models.PostDetailModel, error) {
	post, err := service.GetPostDetailDataById(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) GetListPostDataProcess(postParam models.GetListPostDataParam) ([]*models.PostModel, error) {
	post, err := service.GetListPostDataDataByUserId(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) GetListPostDataDetailProcess(postParam models.GetListPostDataDetailParam) ([]*models.PostDetailModel, error) {
	post, err := service.GetListPostDataDetailByPostId(postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) DeletePostDataProcess(id string, postParam models.DeletePostModel) (interface{}, error) {
	post, err := service.DeletePostData(id, postParam)
	service.DeleteChildRelationData("post_detail", postParam, bson.M{"postid": id})

	if err != nil {
		panic(err)
	}
	return post, nil
}

func (service *PostService) DeletePostDataDetailProcess(id string, postParam models.DeletePostModel) (interface{}, error) {
	post, err := service.DeletePostDataDetail(id, postParam)
	if err != nil {
		panic(err)
	}
	return post, nil
}
