package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginRepositoryWithCircuitBreaker struct {
	LoginRepository interfaces.ILoginRepository
}

func (repository *LoginRepositoryWithCircuitBreaker) GetUserByEmailAndPassword(username string, password string) (models.UserModel, error) {
	output := make(chan models.UserModel, 1)
	hystrix.ConfigureCommand("get_user_by_username_and_password", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_by_username_and_password", func() error {
		user, _ := repository.LoginRepository.GetUserByEmailAndPassword(username, password)
		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.UserModel{}, err
	}
}

type LoginRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *LoginRepository) GetUserByEmailAndPassword(username string, password string) (models.UserModel, error) {
	filter := bson.M{"userName": username, "password": password}
	row, err := repository.FindOne(filter, "users", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return models.UserModel{}, nil
	}

	var user models.UserModel
	row.DecodeResults(&user)
	return user, nil
}
