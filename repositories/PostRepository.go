package repositories

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type PostRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *PostRepository) CreatePost(param models.PostModel) (interface{}, error) {
	fmt.Println(param.UserId)
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
