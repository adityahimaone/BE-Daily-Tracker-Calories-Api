package foods_test

import (
	"daily-tracker-calories/bussiness/foodAPI"
	"daily-tracker-calories/bussiness/foods"
	_mocks "daily-tracker-calories/bussiness/foods/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockFoodRepository _mocks.Repository
	foodService        foods.Service
	foodDomain         foods.Domain
	foodAPIRepo        foodAPI.Repository
)

func TestMain(m *testing.M) {
	foodService = foods.NewService(&mockFoodRepository, foodAPIRepo)
	foodDomain = foods.Domain{
		ID:      1,
		Name:    "food test",
		Calorie: 100,
		Photo:   "photo.jpg",
	}
	m.Run()
}

func TestInsertFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      1,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.SaveFood(&inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.SaveFood(&inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
}

func TestUpdateFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Update", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      1,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.EditFood(inputUser.ID, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Update", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.EditFood(inputUser.ID, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
}

func TestDeleteFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Delete", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      1,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.DeleteFood(inputUser.ID, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Delete", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, err := foodService.DeleteFood(inputUser.ID, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByID", mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByID", mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByID(2)
		assert.Nil(t, err)
		assert.Equal(t, result.ID, foodDomain.ID)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByName("food test")
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByName("food_test")
		assert.Nil(t, err)
		assert.Equal(t, result.ID, foodDomain.ID)
	})
}

func TestGetAllFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetAllFood", mock.Anything, mock.Anything).Return(&[]foods.Domain{foodDomain}, nil).Once()
		result, err := foodService.GetAllFood()
		assert.Nil(t, err)
		assert.Equal(t, &[]foods.Domain{foodDomain}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetAllFood", mock.Anything, mock.Anything).Return(&[]foods.Domain{foodDomain}, nil).Once()
		result, err := foodService.GetAllFood()
		assert.Nil(t, err)
		assert.Equal(t, &[]foods.Domain{foodDomain}, result)
	})
}

// t.Run("Valid Test", func (t *testing.T){})
