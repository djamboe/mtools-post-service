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

func (controller *PostController) createPostDetail(param models.PostDetailModel) (interface{}, error) {
	createPost, err := controller.CreatePostDetailProcess(param)
	if err != nil {
		panic(err)
	}
	return createPost, nil
}

func (controller *PostController) updatePostDetail(id string, param models.PostDetailModel) (interface{}, error) {
	createPost, err := controller.UpdatePostDetailProcess(id, param)
	if err != nil {
		panic(err)
	}
	return createPost, nil
}

func (controller *PostController) postData(param models.PostDataParamModel) (models.PostModel, error) {
	postData, err := controller.GetPostDataProcess(param)
	if err != nil {
		panic(err)
	}
	return postData, nil
}
