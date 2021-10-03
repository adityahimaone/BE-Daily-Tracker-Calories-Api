package histories

import (
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/bussiness/users"
	"log"
	"time"
)

type serviceHistories struct {
	historiesRepository Repository
	foodsRepository     foods.Repository
	usersService        users.Service
	caloriesService     calories.Service
}

func NewService(repositoryHistories Repository, repositoryFoods foods.Repository, serviceUser users.Service, serviceCalorie calories.Service) Service {
	return &serviceHistories{
		historiesRepository: repositoryHistories,
		foodsRepository:     repositoryFoods,
		usersService:        serviceUser,
		caloriesService:     serviceCalorie,
	}
}

func (service *serviceHistories) CreateHistories(histories *Domain) (*Domain, error) {
	user, err := service.usersService.FindByID(histories.UserID)
	if err != nil {
		return &Domain{}, err
	}
	histories.UserID = user.ID
	histories.NameUser = user.Name
	log.Println(user.Name)
	food, err := service.foodsRepository.GetFoodByName(histories.FoodName)
	if err != nil {
		return &Domain{}, err
	}
	histories.FoodID = food.ID
	histories.FoodName = food.Name
	histories.Calorie = food.Calorie
	log.Println(food.Name)
	dateTime := time.Now().Format("2012006")
	histories.Date = dateTime
	log.Println(histories)
	result, err := service.historiesRepository.Insert(histories)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceHistories) GetAllHistoriesByUserID(userid int) (*[]Domain, error) {
	user, err := service.historiesRepository.GetAllHistoriesByUserID(userid)
	if err != nil {
		return &[]Domain{}, err
	}
	return user, nil
}

func (service *serviceHistories) UserStat(userid int) (float64, float64, string, error) {
	currentCalorie, err := service.historiesRepository.SumCalorieByUserID(userid)
	if err != nil {
		return 0.0, 0.0, "", err
	}
	needCalorie, err := service.caloriesService.GetCalorieFloat(userid)
	if err != nil {
		return 0.0, 0.0, "", err
	}
	status := ""
	percent := currentCalorie / needCalorie
	result := percent * 100
	log.Print(result)
	if result < 80 {
		status = "Kurang Makan (<80%)"
	} else if result >= 80 && result <= 100 {
		status = "Cukup Makan (80 - 100%)"
	} else {
		status = "Kelebihan Makan (>100%)"
	}
	return currentCalorie, needCalorie, status, nil
}
