package v1

import (
	"context"
	"fmt"
	"github.com/djamboe/mtools-post-service/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	v1 "github.com/djamboe/mtools-post-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type postServiceServer struct {
}

func NewPostServiceServer() v1.PostServiceServer {
	return &postServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *postServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	} else {
		return status.Errorf(codes.Unimplemented,
			"please input your api version")
	}
	return nil
}

// Create new todo task
func (s *postServiceServer) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully create new post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	photo := make([]models.Photo, len(req.Photo))
	for i, value := range req.Photo {
		photo[i].Id = value.Id
		photo[i].Url = value.Url
	}

	postParam.CustomerId = req.CustomerId
	postParam.CustomerName = req.CustomerName
	postParam.UserId = req.UserId
	postParam.Chanel = req.Chanel
	postParam.Description = req.Description
	postParam.Product = req.Product
	postParam.Phone = req.Phone
	postParam.Pic = req.Pic
	postParam.Price = req.Price
	postParam.Notes = req.Notes
	postParam.Status = req.Status
	postParam.CreatedOn = time.Now()
	postParam.UpdatedOn = time.Now()
	postParam.Photo = photo
	response, err := postController.CreatePostProcess(postParam)

	if err != nil {
		fmt.Println(err)
	}
	if response == 0 {
		message = "Failed insert new post"
	}
	return &v1.CreatePostResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}
