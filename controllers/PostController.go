package controllers

import (
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
)

type PostController struct {
	interfaces.IPostService
}

func (controller *PostController) createPost(param models.PostModel) (interface{}, error) {
	createPost, err := controller.CreatePostProcess(param)
	if err != nil {
		panic(err)
	}
	return createPost, nil
}

func (controller *PostController) updatePost(id string, param models.PostModel) (interface{}, error) {
	createPost, err := controller.UpdatePostProcess(id, param)
	if err != nil {
		panic(err)
	}
	return createPost, nil
}
