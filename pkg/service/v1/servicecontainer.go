package v1

import (
	"context"
	"github.com/djamboe/mtools-post-service/controllers"
	"github.com/djamboe/mtools-post-service/infrastructures"
	"github.com/djamboe/mtools-post-service/repositories"
	"github.com/djamboe/mtools-post-service/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type IserviceContainer interface {
	InjectPostController() controllers.PostController
}

type kernel struct{}

func (k *kernel) InjectPostController() controllers.PostController {
	//mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/marketing-tools?charset=utf8")
	//mysqlHandler := &infrastructures.DBHandler{}
	//mysqlHandler.Conn = mysqlConn
	//loginRepository := &repositories.LoginRepository{mysqlHandler}
	//loginService := &services.LoginService{&repositories.LoginRepositoryWithCircuitBreaker{loginRepository}}
	//loginController := controllers.LoginController{loginService}

	//test mongodb
	c := GetClient()

	mongoDBConn := c
	mongoDBHandler := &infrastructures.MongoDBHandler{}
	mongoDBHandler.Conn = mongoDBConn
	postRepository := &repositories.PostRepository{mongoDBHandler}
	postService := &services.PostService{&repositories.PostRepositoryWithCircuitBreaker{postRepository}}
	postController := controllers.PostController{postService}
	//test mongodb

	return postController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IserviceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
