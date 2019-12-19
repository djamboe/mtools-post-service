package v1

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/djamboe/mtools-post-service/models"
	v1 "github.com/djamboe/mtools-post-service/pkg/api/v1"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	//v1 "github.com/djamboe/mtools-post-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type postServiceServer struct {
}

var conn *amqp.Connection
var ch *amqp.Channel

func initAmqp() {
	var err error

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"post-data-exchange", // name
		"direct",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare the Exchange")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
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

// Create new Post
func (s *postServiceServer) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	flag.Parse()
	initAmqp()
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully create new post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

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
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"post-data-exchange", // exchange
		"go-test-key",        // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		fmt.Println(err)
	}

	return &v1.CreatePostResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	flag.Parse()
	initAmqp()
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully update post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	photo := make([]models.Photo, len(req.Photo))
	for i, value := range req.Photo {
		photo[i].Id = value.Id
		photo[i].Url = value.Url
	}
	postParam.DbId = req.DbId
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
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"update-post-exchange", // exchange
		"go-test-key-update",   // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		fmt.Println(err)
	}

	return &v1.UpdatePostResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) CreatePostDetail(ctx context.Context, req *v1.CreatePostDetailRequest) (*v1.CreatePostDetailResponse, error) {
	flag.Parse()
	initAmqp()
	var postParam models.PostDetailParamModel
	// check if the API version requested by client is supported by server
	message := "Successfully create new post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	photo := make([]models.Photo, len(req.Photo))
	for i, value := range req.Photo {
		photo[i].Id = value.Id
		photo[i].Url = value.Url
	}

	postParam.PostId = req.PostId
	postParam.Notes = req.Notes
	postParam.Description = req.Description
	postParam.Description = req.Description
	postParam.Notes = req.Notes
	postParam.Status = req.Status
	postParam.CreatedOn = time.Now()
	postParam.UpdatedOn = time.Now()
	postParam.IsDeleted = false
	postParam.Photo = photo
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"post-detail-data-exchange", // exchange
		"go-test-key",               // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		fmt.Println(err)
	}

	return &v1.CreatePostDetailResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) UpdatePostDetail(ctx context.Context, req *v1.UpdatePostDetailRequest) (*v1.UpdatePostDetailResponse, error) {
	flag.Parse()
	initAmqp()
	var postParam models.PostDetailParamModel
	// check if the API version requested by client is supported by server
	message := "Successfully update post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	photo := make([]models.Photo, len(req.Photo))
	for i, value := range req.Photo {
		photo[i].Id = value.Id
		photo[i].Url = value.Url
	}
	postParam.DbId = req.DbId
	postParam.Description = req.Description
	postParam.Notes = req.Notes
	postParam.Status = req.Status
	postParam.CreatedOn = time.Now()
	postParam.UpdatedOn = time.Now()
	postParam.Photo = photo
	postParam.IsDeleted = req.IsDeleted
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"update-post-detail-exchange", // exchange
		"go-test-key-update-detail",   // routing key
		false,                         // mandatory
		false,                         // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		fmt.Println(err)
	}

	return &v1.UpdatePostDetailResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) GetPostData(ctx context.Context, req *v1.GetPostDataRequest) (*v1.GetPostDataResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully get post data !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	postData := models.PostDataParamModel{}
	postData.Id = req.Id
	response, err := postController.GetPostDataProcess(postData)

	if err != nil {
		message = "Failed get post data !"
		errorStatus = true
	}

	postDataResponse := &v1.Post{}
	postDataResponse.UserId = response.UserId
	postDataResponse.Notes = response.Notes
	postDataResponse.Phone = response.Phone
	postDataResponse.Price = response.Price
	postDataResponse.Pic = response.Pic
	postDataResponse.Product = response.Product
	postDataResponse.Status = response.Status
	postDataResponse.CustomerId = response.CustomerId
	postDataResponse.CustomerName = response.CustomerName
	postDataResponse.Description = response.Description
	postDataResponse.Chanel = response.Chanel
	postDataResponse.IsDelete = response.IsDeleted
	return &v1.GetPostDataResponse{
		Api:     apiVersion,
		Error:   errorStatus,
		Message: message,
		Post:    postDataResponse,
	}, nil
}

func (s *postServiceServer) GetPostDataDetail(ctx context.Context, req *v1.GetPostDataDetailRequest) (*v1.GetPostDataDetailResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully get post detail data !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	postData := models.GetPostDetailParamModel{}
	postData.Id = req.Id
	response, err := postController.GetPostDetailDataProcess(postData)

	if err != nil {
		message = "Failed get post data !"
		errorStatus = true
	}

	postDataResponse := &v1.PostDetail{}
	postDataResponse.PostId = response.PostId
	postDataResponse.Notes = response.Notes
	postDataResponse.Status = response.Status
	postDataResponse.Description = response.Description
	postDataResponse.IsDelete = response.IsDeleted
	postDataResponse.PostId = response.PostId

	return &v1.GetPostDataDetailResponse{
		Api:        apiVersion,
		Error:      errorStatus,
		Message:    message,
		PostDetail: postDataResponse,
	}, nil
}
