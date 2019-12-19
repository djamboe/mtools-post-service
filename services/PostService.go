package services

import (
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
)

type PostService struct {
	interfaces.IPostRepository
}

func (service *PostService) CreatePostProcess(postParam models.PostModel) (interface{}, error) {
	user, err := service.CreatePost(postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) UpdatePostProcess(id string, postParam models.PostModel) (interface{}, error) {
	user, err := service.UpdatePost(id, postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) CreatePostDetailProcess(postParam models.PostDetailModel) (interface{}, error) {
	user, err := service.CreatePostDetail(postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) UpdatePostDetailProcess(id string, postParam models.PostDetailModel) (interface{}, error) {
	user, err := service.UpdatePostDetail(id, postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) GetPostDetailData(id string, postParam models.PostDetailModel) (interface{}, error) {
	user, err := service.UpdatePostDetail(id, postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) GetPostDataProcess(postParam models.PostDataParamModel) (models.PostModel, error) {
	user, err := service.GetPostDataById(postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *PostService) GetPostDetailDataProcess(postParam models.GetPostDetailParamModel) (models.PostDetailModel, error) {
	user, err := service.GetPostDetailDataById(postParam)
	if err != nil {
		panic(err)
	}
	return user, nil
}
