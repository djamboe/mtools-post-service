package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/djamboe/mtools-post-service/models"
	v1 "github.com/djamboe/mtools-post-service/pkg/api/v1"
	structpb "github.com/golang/protobuf/ptypes/struct"
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

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type postServiceServer struct {
}

func init() {
	initAmqp()
}

var conn *amqp.Connection
var ch *amqp.Channel

func initAmqp() {
	var err error

	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
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
	var postParam models.PostModelParam
	var loc, _ = time.LoadLocation("Asia/Jakarta")
	// check if the API version requested by client is supported by server
	message := "Successfully create new post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	postParam.CustomerId = req.Customerid
	postParam.CustomerName = req.Customername
	postParam.UserId = req.Userid
	postParam.Chanel = req.Channel
	postParam.Description = req.Description
	postParam.Product = req.Product
	postParam.ProductId = req.Productid
	postParam.Phone = req.Phone
	postParam.Pic = req.Pic
	postParam.Price = req.Price
	postParam.Notes = req.Notes
	postParam.Status = req.Status
	postParam.CreatedOn = time.Now().In(loc)
	postParam.UpdatedOn = time.Now().In(loc)
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
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully update post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	postParam.DbId = req.DbId
	postParam.CustomerId = req.Customerid
	postParam.CustomerName = req.Customername
	postParam.ProductId = req.Productid
	postParam.UserId = req.Userid
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
	postDataResponse.Userid = response.UserId
	postDataResponse.Notes = response.Notes
	postDataResponse.Phone = response.Phone
	postDataResponse.Price = response.Price
	postDataResponse.Pic = response.Pic
	postDataResponse.Productid = response.ProductId
	postDataResponse.Product = response.Product
	postDataResponse.Status = response.Status
	postDataResponse.Customerid = response.CustomerId
	postDataResponse.Customername = response.CustomerName
	postDataResponse.Description = response.Description
	postDataResponse.Channel = response.Chanel
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
	photo := new(v1.Photo)
	for _, value := range response.Photo {
		photo.Id = value.Id
		photo.Url = value.Url
		postDataResponse.Photo = append(postDataResponse.Photo, photo)
	}

	return &v1.GetPostDataDetailResponse{
		Api:        apiVersion,
		Error:      errorStatus,
		Message:    message,
		PostDetail: postDataResponse,
	}, nil
}

func (s *postServiceServer) GetListPostData(ctx context.Context, req *v1.GetListPostDataRequest) (*v1.GetListPostDataResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully get list post data !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	listPostData := models.GetListPostDataParam{}
	listPostData.UserId = req.UserId
	response, err := postController.GetListPostDataProcess(listPostData)

	if err != nil {
		message = "Failed get list post data !"
		errorStatus = true
	}

	postDataSlice := make([]*v1.Post, len(response))
	//photo := new(v1.Photo)
	for i, value := range response {
		postDataSlice[i] = new(v1.Post)
		postDataSlice[i].DbId = value.DbId.Hex()
		postDataSlice[i].Userid = value.UserId
		postDataSlice[i].Customerid = value.CustomerId
		postDataSlice[i].Customername = value.CustomerName
		postDataSlice[i].Userid = value.UserId
		postDataSlice[i].Productid = value.ProductId
		postDataSlice[i].Channel = value.Chanel
		postDataSlice[i].Description = value.Description
		postDataSlice[i].Product = value.Product
		postDataSlice[i].Phone = value.Phone
		postDataSlice[i].CreatedOn = value.CreatedOn.Format("2006-01-02 15:04:05")
		postDataSlice[i].Pic = value.Pic
		postDataSlice[i].Price = value.Price
		postDataSlice[i].Notes = value.Notes
		postDataSlice[i].Status = value.Status
	}

	return &v1.GetListPostDataResponse{
		Api:     apiVersion,
		Error:   errorStatus,
		Message: message,
		Post:    postDataSlice,
	}, nil
}

func (s *postServiceServer) GetListPostDataDetail(ctx context.Context, req *v1.GetListPostDataDetailRequest) (*v1.GetListPostDataDetailResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully get list post data !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	listPostDataDetail := models.GetListPostDataDetailParam{}
	listPostDataDetail.PostId = req.PostId
	response, err := postController.GetListPostDataDetailProcess(listPostDataDetail)

	if err != nil {
		message = "Failed get list post data !"
		errorStatus = true
	}

	postDataDetailSlice := make([]*v1.PostDetail, len(response))
	//photo := new(v1.Photo)
	for i, value := range response {
		postDataDetailSlice[i] = new(v1.PostDetail)
		postDataDetailSlice[i].PostId = value.PostId
		postDataDetailSlice[i].Description = value.Description
		postDataDetailSlice[i].Notes = value.Notes
		postDataDetailSlice[i].Status = value.Status
		photo := make([]*v1.Photo, len(value.Photo))
		for i, valuePhoto := range value.Photo {
			photo[i] = new(v1.Photo)
			photo[i].Id = valuePhoto.Id
			photo[i].Url = valuePhoto.Url
		}
		postDataDetailSlice[i].Photo = photo
	}

	return &v1.GetListPostDataDetailResponse{
		Api:        apiVersion,
		Error:      errorStatus,
		Message:    message,
		PostDetail: postDataDetailSlice,
	}, nil
}

func (s *postServiceServer) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully deleted post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	postParam.DbId = req.Id
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"delete-post-exchange", // exchange
		"go-test-key-delete",   // routing key
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

	return &v1.DeletePostResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) DeletePostDetail(ctx context.Context, req *v1.DeletePostDetailRequest) (*v1.DeletePostDetailResponse, error) {
	var postParam models.PostModelParam
	// check if the API version requested by client is supported by server
	message := "Successfully deleted post !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}

	postParam.DbId = req.Id
	payload, err := json.Marshal(postParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"delete-post-detail-exchange", // exchange
		"go-test-key-delete-detail",   // routing key
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

	return &v1.DeletePostDetailResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}

func (s *postServiceServer) GetWeeklyPlanData(ctx context.Context, req *v1.GetWeeklyPlanRequest) (*v1.GetWeeklyPlanResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully get list weekly plan data !"
	errorStatus := false

	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	postController := ServiceContainer().InjectPostController()
	listWeeklyPlanData := models.GetWeeklyPlanParamModel{}
	listWeeklyPlanData.Month = req.Month
	response, err := postController.GetListWeeklyPlanDataProcess(listWeeklyPlanData)

	if err != nil {
		message = "Failed get list weekly plan data !"
		errorStatus = true
	}

	x := make(map[string][]string)
	for _, value := range response {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		weeklyPlanDate := value.Date.Time().UTC().In(loc)
		formattedDate := fmt.Sprintf("%d-%02d-%02d",
			weeklyPlanDate.Year(), weeklyPlanDate.Month(), weeklyPlanDate.Day())

		x[formattedDate] = append(x[formattedDate], value.Title+"|"+value.DbId.Hex())
	}

	weeklyPlanDataSlice := make([]*v1.Plan, len(x))
	if len(x) > 0 {
		iterator := 0
		for key, value := range x {
			weeklyPlanDataSlice[iterator] = new(v1.Plan)
			weeklyPlanDataSlice[iterator].Date = key

			planTitle := new(structpb.ListValue)
			for _, valueTitle := range value {
				planTitle.Values = append(planTitle.Values, &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: valueTitle}})
			}
			weeklyPlanDataSlice[iterator].PlanTitle = planTitle
			iterator++
		}
	}

	return &v1.GetWeeklyPlanResponse{
		Api:     apiVersion,
		Error:   errorStatus,
		Message: message,
		Plan:    weeklyPlanDataSlice,
	}, nil
}
