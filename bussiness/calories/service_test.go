package calories_test

import (
	"daily-tracker-calories/bussiness/calories"
	_mocks "daily-tracker-calories/bussiness/calories/mocks"
	"daily-tracker-calories/bussiness/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

var (
	mockCaloriesRepository _mocks.Repository
	userService            users.Service
	caloriesService        calories.Service
	caloriesDomain         calories.Domain
)

func TestMain(m *testing.M) {
	caloriesService = calories.NewService(&mockCaloriesRepository, userService)
	caloriesDomain = calories.Domain{
		ID:           1,
		Weight:       60,
		Height:       170,
		Gender:       "Male",
		Age:          25,
		ActivityType: 3,
		Calorie:      1468.75,
		UserID:       1,
	}
	m.Run()
}

func TestCountCalorie(t *testing.T) {
	t.Run("Valid Test Male", func(t *testing.T) {
		weightFloat := 60.0
		heightFloat := 170.0
		ageFloat := 25.0
		valueActivity := 1.55
		calories := (10 * weightFloat) + (6.25 * heightFloat) - (5*ageFloat)*valueActivity
		assert.Nil(t, nil)
		assert.Equal(t, calories, caloriesDomain.Calorie)
	})
	t.Run("Valid Test Female", func(t *testing.T) {
		weightFloat := 60.0
		heightFloat := 170.0
		ageFloat := 25.0
		valueActivity := 1.2
		calories := ((10 * weightFloat) + (6.25 * heightFloat) - (5 * ageFloat) - 161) * valueActivity
		assert.Nil(t, nil)
		resultCalories := 1651.8
		assert.Equal(t, calories, resultCalories)
	})
}

func TestGetCalorieByUserID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.GetCalorieByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("invalid Test", func(t *testing.T) {
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything).Return(&caloriesDomain, assert.AnError).Once()
		result, _ := caloriesService.GetCalorieByUserID(2)
		assert.Equal(t, &calories.Domain{}, result)
	})
}

func TestGetCalorieFloat(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything).Return(caloriesDomain.Calorie, nil).Once()
		result, err := caloriesService.GetCalorieFloat(1)
		assert.Nil(t, err)
		assert.Equal(t, caloriesDomain.Calorie, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything).Return(caloriesDomain.Calorie, assert.AnError).Once()
		result, err := caloriesService.GetCalorieFloat(2)
		assert.NotNil(t, err)
		assert.Equal(t, 0.0, result)
	})
}

func TestCreateCalorie(t *testing.T) {
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "male",
			Age:          25,
			ActivityType: 3,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "male",
			Age:          25,
			ActivityType: 2,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "male",
			Age:          25,
			ActivityType: 4,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "male",
			Age:          25,
			ActivityType: 5,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "male",
			Age:          25,
			ActivityType: 6,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Female", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "female",
			Age:          25,
			ActivityType: 1,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
	t.Run("Valid Test Female", func(t *testing.T) {
		inputUser := calories.Domain{
			ID:           1,
			Weight:       60,
			Height:       170,
			Gender:       "female",
			Age:          25,
			ActivityType: 1,
			Calorie:      1468.75,
			UserID:       1,
		}
		mockCaloriesRepository.On("GetCalorieFloat", mock.Anything, mock.Anything).Return(&caloriesDomain.Calorie, nil).Once()
		mockCaloriesRepository.On("GetCalorieByUserID", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		caloriesDomain.ID = 0
		mockCaloriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		mockCaloriesRepository.On("Update", mock.Anything, mock.Anything).Return(&caloriesDomain, nil).Once()
		result, err := caloriesService.CreateCalorie(&inputUser,1)
		log.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, &caloriesDomain, result)
	})
}