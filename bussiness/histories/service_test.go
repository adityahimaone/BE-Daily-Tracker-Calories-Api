package histories_test

import (
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/bussiness/histories"
	_mocks "daily-tracker-calories/bussiness/histories/mocks"
	"daily-tracker-calories/bussiness/users"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockHistoriesRepository _mocks.Repository
	historiesService        histories.Service
	historiesDomain         histories.Domain
	userService             users.Service
	foodService             foods.Service
	foodRepository          foods.Repository
	calorieService          calories.Service
)

func TestMain(m *testing.M) {
	historiesService = histories.NewService(&mockHistoriesRepository, foodRepository, userService, calorieService, foodService)
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
		mockHistoriesRepository.On("GetAllHistoriesByUserID", mock.Anything, mock.Anything).Return(&[]histories.Domain{}, errors.New("hello world")).Once()
		result, _ := historiesService.GetAllHistoriesByUserID(1)
		assert.Equal(t, &[]histories.Domain{}, result)
	})
}
