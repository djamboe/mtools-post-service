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
