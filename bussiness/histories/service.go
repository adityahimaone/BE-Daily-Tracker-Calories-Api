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

func (s *serviceHistories) CreateHistories(histories *Domain) (*Domain, error) {
	user, err := s.usersService.FindByID(histories.UserID)
	if err != nil {
		return &Domain{}, err
	}
	histories.UserID = user.ID
	histories.NameUser = user.Name
	log.Println(user.Name)
	food, err := s.foodsRepository.GetFoodByName(histories.FoodName)
	if err != nil {
		return &Domain{}, err
	}
	histories.FoodID = food.ID
	histories.FoodName = food.Name
	histories.Calorie = food.Calorie
	log.Println(food.Name)
	dateTime := time.Now()
	histories.Date = dateTime
	log.Println(histories)
	result, err := s.historiesRepository.Insert(histories, food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceHistories) GetHistoriesByUserID(userid int) (*[]Domain, error) {
	panic("implement me")
}
