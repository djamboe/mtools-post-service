package controllers

import (
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
)

type PostController struct {
	interfaces.IPostService
}

func (controller *PostController) createPost(param models.PostModelParam) (interface{}, error) {
	createPost, err := controller.CreatePostProcess(param)
	if err != nil {
		panic(err)
	}
	return createPost, nil
}
