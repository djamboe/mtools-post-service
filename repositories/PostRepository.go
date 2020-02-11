package repositories

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-post-service/interfaces"
	"github.com/djamboe/mtools-post-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
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

func (repository *PostRepositoryWithCircuitBreaker) GetListWeeklyPlanData(param models.GetWeeklyPlanParamModel) ([]*models.WeeklyPlan, error) {
	output := make(chan []*models.WeeklyPlan, 1)
	hystrix.ConfigureCommand("list_weekly_plan", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("list_weekly_plan", func() error {
		weeklyPlanData, _ := repository.PostRepository.GetListWeeklyPlanData(param)
		output <- weeklyPlanData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []*models.WeeklyPlan{}, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) GetListPostDataDataByUserId(param models.GetListPostDataParam) ([]*models.PostModel, error) {
	output := make(chan []*models.PostModel, 1)
	hystrix.ConfigureCommand("list_post_data", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("list_post_data", func() error {
		postData, _ := repository.PostRepository.GetListPostDataDataByUserId(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []*models.PostModel{}, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) GetListPostDataDetailByPostId(param models.GetListPostDataDetailParam) ([]*models.PostDetailModel, error) {
	output := make(chan []*models.PostDetailModel, 1)
	hystrix.ConfigureCommand("list_post_data_detail", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("list_post_data_detail", func() error {
		postData, _ := repository.PostRepository.GetListPostDataDetailByPostId(param)
		output <- postData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []*models.PostDetailModel{}, err
	}
}

func (repository *PostRepositoryWithCircuitBreaker) DeletePostData(id string, param models.DeletePostModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("delete_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("delete_post", func() error {
		postData, _ := repository.PostRepository.DeletePostData(id, param)
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

func (repository *PostRepositoryWithCircuitBreaker) DeletePostDataDetail(id string, param models.DeletePostModel) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("delete_post_detail", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("delete_post_detail", func() error {
		postData, _ := repository.PostRepository.DeletePostDataDetail(id, param)
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

func (repository *PostRepositoryWithCircuitBreaker) DeleteChildRelationData(collectionName string, param models.DeletePostModel, filter bson.M) (interface{}, error) {
	output := make(chan interface{}, 1)
	hystrix.ConfigureCommand("update_child_relation", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("update_child_relation", func() error {
		postData, _ := repository.PostRepository.DeleteChildRelationData(collectionName, param, filter)
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

func (repository *PostRepository) GetListPostDataDataByUserId(dataPostParamModels models.GetListPostDataParam) ([]*models.PostModel, error) {
	docId := dataPostParamModels.UserId
	filter := bson.M{"userid": docId}
	row, err := repository.Find(filter, "post", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return []*models.PostModel{}, nil
	}

	var listPostData []*models.PostModel
	for row.Next(context.TODO()) {
		var data models.PostModel
		err = row.Decode(&data)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listPostData = append(listPostData, &data)
	}
	return listPostData, nil
}

func (repository *PostRepository) GetListPostDataDetailByPostId(dataPostParamModels models.GetListPostDataDetailParam) ([]*models.PostDetailModel, error) {
	docId := dataPostParamModels.PostId
	filter := bson.M{"postid": docId}
	row, err := repository.Find(filter, "post_detail", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return []*models.PostDetailModel{}, nil
	}

	var listPostData []*models.PostDetailModel
	for row.Next(context.TODO()) {
		var data models.PostDetailModel
		err = row.Decode(&data)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listPostData = append(listPostData, &data)
	}
	return listPostData, nil
}

func (repository *PostRepository) DeletePostData(id string, param models.DeletePostModel) (interface{}, error) {
	row, err := repository.UpdateOne(id, param, "post", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) DeletePostDataDetail(id string, param models.DeletePostModel) (interface{}, error) {
	row, err := repository.UpdateOne(id, param, "post_detail", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) DeleteChildRelationData(collectionName string, param models.DeletePostModel, filter bson.M) (interface{}, error) {
	row, err := repository.UpdateMany(param, collectionName, "maroon_martools", filter)

	if err != nil {
		panic(err)
	}

	if row == nil {
		return 0, nil
	}

	return row, nil
}

func (repository *PostRepository) GetListWeeklyPlanData(dataWeeklyPlanParamModel models.GetWeeklyPlanParamModel) ([]*models.WeeklyPlan, error) {
	currentTime := time.Now()

	currentMonth := time.January

	if dataWeeklyPlanParamModel.Month == "01" {
		currentMonth = time.January
	} else if dataWeeklyPlanParamModel.Month == "02" {
		currentMonth = time.February
	} else if dataWeeklyPlanParamModel.Month == "03" {
		currentMonth = time.March
	} else if dataWeeklyPlanParamModel.Month == "04" {
		currentMonth = time.April
	} else if dataWeeklyPlanParamModel.Month == "05" {
		currentMonth = time.May
	} else if dataWeeklyPlanParamModel.Month == "06" {
		currentMonth = time.June
	} else if dataWeeklyPlanParamModel.Month == "07" {
		currentMonth = time.July
	} else if dataWeeklyPlanParamModel.Month == "08" {
		currentMonth = time.August
	} else if dataWeeklyPlanParamModel.Month == "09" {
		currentMonth = time.September
	} else if dataWeeklyPlanParamModel.Month == "10" {
		currentMonth = time.October
	} else if dataWeeklyPlanParamModel.Month == "11" {
		currentMonth = time.November
	} else if dataWeeklyPlanParamModel.Month == "12" {
		currentMonth = time.December
	}
	//month := dataWeeklyPlanParamModel.Month
	fromDate := time.Date(currentTime.Year(), currentMonth, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(currentTime.Year(), currentMonth, 31, 0, 0, 0, 0, time.UTC)
	filter := bson.M{"date": bson.M{
		"$gt":  fromDate,
		"$lte": toDate,
	}}
	//filter := bson.M{"userid": docId}
	row, err := repository.Find(filter, "weekly_plan", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return []*models.WeeklyPlan{}, nil
	}

	var listWeeklyPlanData []*models.WeeklyPlan
	for row.Next(context.TODO()) {
		var data models.WeeklyPlan
		err = row.Decode(&data)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listWeeklyPlanData = append(listWeeklyPlanData, &data)
	}
	return listWeeklyPlanData, nil
}
