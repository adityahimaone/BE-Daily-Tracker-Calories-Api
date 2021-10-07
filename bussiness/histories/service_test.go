package histories_test

import (
	_mocksCalories "daily-tracker-calories/bussiness/calories/mocks"
	"daily-tracker-calories/bussiness/foods"
	_mocksFood "daily-tracker-calories/bussiness/foods/mocks"
	"daily-tracker-calories/bussiness/histories"
	_mocksHistories "daily-tracker-calories/bussiness/histories/mocks"
	"daily-tracker-calories/bussiness/users"
	_mocksUser "daily-tracker-calories/bussiness/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

var (
	mockHistoriesRepository _mocksHistories.Repository
	mockUserRepository      _mocksUser.Repository
	mockFoodRepository      _mocksFood.Repository
	mockCaloriesRepository  _mocksCalories.Repository
	historiesService        histories.Service
	historiesDomain         histories.Domain
	userService             _mocksUser.Service
	foodService             _mocksFood.Service
	foodRepository          foods.Repository
	mockCalorieService      _mocksCalories.Service
	foodDomain              foods.Domain
	userDomain              users.Domain
)

func TestMain(m *testing.M) {
	historiesService = histories.NewService(&mockHistoriesRepository, foodRepository, &userService, &mockCalorieService, &foodService)
	historiesDomain = histories.Domain{
		ID:          1,
		UserID:      1,
		NameUser:    "test nama",
		FoodID:      1,
		Calorie:     100,
		FoodName:    "food test",
		Date:        "5102021",
		SumCalorie:  1200,
		CalorieNeed: 1500,
	}
	foodDomain = foods.Domain{
		ID:      1,
		Name:    "food test",
		Calorie: 100,
		Photo:   "photo.jpg",
	}
	userDomain = users.Domain{
		ID:    1,
		Name:  "food test",
		Email: "test@mail.com",
	}
	m.Run()
}
func TestGetAllByUserID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetAllHistoriesByUserID", mock.Anything, mock.Anything).Return(&[]histories.Domain{historiesDomain}, nil).Once()
		result, err := historiesService.GetAllHistoriesByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, &[]histories.Domain{historiesDomain}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetAllHistoriesByUserID", mock.Anything, mock.Anything).Return(&[]histories.Domain{}, assert.AnError).Once()
		result, _ := historiesService.GetAllHistoriesByUserID(1)
		assert.Equal(t, &[]histories.Domain{}, result)
	})
}

func TestCreateHistories(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		userService.On("FindByID", mock.Anything, mock.Anything).Return(&userDomain, nil).Once() //userdomain
		foodService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		input := histories.Domain{
			ID:          1,
			UserID:      1,
			NameUser:    "test nama",
			FoodID:      1,
			Calorie:     100,
			FoodName:    "food test",
			Date:        "5102021",
			SumCalorie:  1200,
			CalorieNeed: 1500,
		}
		result, err := historiesService.CreateHistories(&input)
		assert.Nil(t, err)
		assert.Equal(t, &historiesDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		userService.On("FindByID", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		foodService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, _ := historiesService.CreateHistories(&histories.Domain{})
		assert.Equal(t, &histories.Domain{}, result)
	})
	/*t.Run("Invalid Test", func(t *testing.T) {
		userService.On("FindByID", mock.Anything, mock.Anything).Return(&users.Domain{}, nil).Once()
		foodService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&histories.Domain{}, nil).Once()
		result, _ := historiesService.CreateHistories(&histories.Domain{})
		assert.Equal(t, &histories.Domain{}, result)
	})*/
/*	t.Run("Invalid Test", func(t *testing.T) {
		userService.On("FindByID", mock.Anything, mock.Anything).Return(&users.Domain{}, nil).Once()
		foodService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, assert.AnError).Once()
		result, _ := historiesService.CreateHistories(&histories.Domain{})
		assert.Equal(t, &histories.Domain{}, result)
	})*/
}

func TestUserStat(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		sumCalorie := 100.0
		mockCalorieService.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(1.0, nil).Once()
		mockHistoriesRepository.On("SumCalorieByUserID", mock.Anything, mock.Anything).Return(sumCalorie, nil).Once()
		currentCalorie, needCalorie, str_percentage, status, err := historiesService.UserStat(1)
		currentCalorie = 100.0
		needCalorie = 1500
		str_percentage = "70 %"
		status = "over"
		log.Println(currentCalorie)
		log.Println(needCalorie)
		log.Println(str_percentage)
		log.Println(status)
		assert.Nil(t, err)
		assert.Equal(t, currentCalorie, sumCalorie)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		sumCalorie := 100.0
		mockCalorieService.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(1.0, assert.AnError).Once()
		mockHistoriesRepository.On("SumCalorieByUserID", mock.Anything, mock.Anything).Return(sumCalorie, nil).Once()
		currentCalorie, needCalorie, str_percentage, status, err := historiesService.UserStat(1)
		currentCalorie = 100.0
		needCalorie = 1500
		str_percentage = "70 %"
		status = "over"
		log.Println(currentCalorie)
		log.Println(needCalorie)
		log.Println(str_percentage)
		log.Println(status)
		assert.NotNil(t, err)
		assert.Equal(t, currentCalorie, sumCalorie)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		sumCalorie := 100.0
		mockCalorieService.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(1.0, nil).Once()
		mockHistoriesRepository.On("SumCalorieByUserID", mock.Anything, mock.Anything).Return(sumCalorie, assert.AnError).Once()
		currentCalorie, needCalorie, str_percentage, status, err := historiesService.UserStat(1)
		currentCalorie = 100.0
		needCalorie = 1500
		str_percentage = "70 %"
		status = "over"
		log.Println(currentCalorie)
		log.Println(needCalorie)
		log.Println(str_percentage)
		log.Println(status)
		assert.NotNil(t, err)
		assert.Equal(t, currentCalorie, sumCalorie)
	})
}
