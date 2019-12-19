package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostRepositoryWithCircuitBreaker struct {
	PostRepository interfaces.IPostRepository
}

func (repository *PostRepositoryWithCircuitBreaker) CreatePost(param models.PostModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("create_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_post", func() error {
		postData, _ := repository.PostRepository.CreatePost(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return 0, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) UpdatePost(id string, param models.PostModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("update_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("update_post", func() error {
		postData, _ := repository.PostRepository.UpdatePost(id, param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return 0, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) CreatePostDetail(param models.PostDetailModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("create_post_detail", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_post_detail", func() error {
		postData, _ := repository.PostRepository.CreatePostDetail(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return 0, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) UpdatePostDetail(id string, param models.PostDetailModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("update_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("update_post", func() error {
		postData, _ := repository.PostRepository.UpdatePostDetail(id, param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return 0, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) GetPostDataById(param models.PostDataParamModel) (models.PostModel, error) {
	output := make(chan models.PostModel, 1)
	hystrix.ConfigureCommand("post_data", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("post_data", func() error {
		postData, _ := repository.PostRepository.GetPostDataById(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.PostModel{}, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) GetPostDetailDataById(param models.GetPostDetailParamModel) (models.PostDetailModel, error) {
	output := make(chan models.PostDetailModel, 1)
	hystrix.ConfigureCommand("post_detail_data", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("post_detail_data", func() error {
		postData, _ := repository.PostRepository.GetPostDetailDataById(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.PostDetailModel{}, err
	}
}

type PostRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *PostRepository) CreatePost(param models.PostModel) (interface{}, error) {
	row, err := repository.InsertOne(param, "post", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) UpdatePost(id string, param models.PostModel) (interface{}, error) {
	row, err := repository.UpdateOne(id, param, "post", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) CreatePostDetail(param models.PostDetailModel) (interface{}, error) {
	row, err := repository.InsertOne(param, "post_detail", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) UpdatePostDetail(id string, param models.PostDetailModel) (interface{}, error) {
	row, err := repository.UpdateOne(id, param, "post_detail", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) GetPostDataById(dataPostParamModels models.PostDataParamModel) (models.PostModel, error) {
	docId := dataPostParamModels.Id
	objId, err := primitive.ObjectIDFromHex(docId)

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	row, err := repository.FindOne(filter, "post", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return models.PostModel{}, nil
	}

	var postData models.PostModel
	row.DecodeResults(&postData)
	return postData, nil
}

func (repository *PostRepository) GetPostDetailDataById(dataPostParamModels models.GetPostDetailParamModel) (models.PostDetailModel, error) {
	docId := dataPostParamModels.Id
	objId, err := primitive.ObjectIDFromHex(docId)

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	row, err := repository.FindOne(filter, "post_detail", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return models.PostDetailModel{}, nil
	}

	var postData models.PostDetailModel
	row.DecodeResults(&postData)
	return postData, nil
}
