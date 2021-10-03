package histories

import (
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/bussiness/users"
	"log"
	"time"
)

type serviceHistories struct {
	historiesRepository Repository
	foodsRepository     foods.Repository
	usersService        users.Service
}

func NewService(repositoryHistories Repository, repositoryFoods foods.Repository, serviceUser users.Service) Service {
	return &serviceHistories{
		historiesRepository: repositoryHistories,
		foodsRepository:     repositoryFoods,
		usersService:        serviceUser,
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
