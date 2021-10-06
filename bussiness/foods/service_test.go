package foods_test

import (
	"daily-tracker-calories/bussiness/foodAPI"
	_mocksFoodAPIRepo "daily-tracker-calories/bussiness/foodAPI/mocks"
	"daily-tracker-calories/bussiness/foods"
	_mocksFoodRepo "daily-tracker-calories/bussiness/foods/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

var (
	mockFoodRepository _mocksFoodRepo.Repository
	foodService        foods.Service
	foodDomain         foods.Domain
	foodAPIDomain      foodAPI.Domain
	mockFoodAPIRepo    _mocksFoodAPIRepo.Repository
)

func TestMain(m *testing.M) {
	foodService = foods.NewService(&mockFoodRepository, &mockFoodAPIRepo)
	foodDomain = foods.Domain{
		/*ID:      1,*/
		Name:    "food test",
		Calorie: 100,
		Photo:   "photo.jpg",
	}
	foodAPIDomain = foodAPI.Domain{
		Name:    "food test",
		Photo:   "photo.jpg",
		Calorie: 100,
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
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodDomain, assert.AnError).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, _ := foodService.SaveFood(&inputUser)
		assert.Equal(t, &foods.Domain{}, result)
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
		result, err := foodService.EditFood(1, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Update", mock.Anything, mock.Anything).Return(&foodDomain, assert.AnError).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, _ := foodService.EditFood(2, &inputUser)
		assert.Equal(t, &foods.Domain{}, result)
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
		result, err := foodService.DeleteFood(1, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Delete", mock.Anything, mock.Anything).Return(&foodDomain, assert.AnError).Once()
		inputUser := foods.Domain{
			ID:      2,
			Name:    "food test",
			Calorie: 100,
			Photo:   "photo.jpg",
		}
		result, _ := foodService.DeleteFood(2, &inputUser)
		assert.Equal(t, &foods.Domain{}, result)
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
		mockFoodRepository.On("GetFoodByID", mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		result, _ := foodService.GetFoodByID(2)
		log.Println(result)
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByName("food test")
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		mockFoodAPIRepo.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodAPIDomain, nil).Once()
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		result, err := foodService.GetFoodByName("food test")
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		mockFoodAPIRepo.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodAPIDomain, assert.AnError).Once()
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodDomain, nil).Once()
		result, _ := foodService.GetFoodByName("food test")
		assert.Equal(t, &foods.Domain{}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		mockFoodAPIRepo.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodAPI.Domain{}, assert.AnError).Once()
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		result, _ := foodService.GetFoodByName("bukanmakanan")
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestGetFoodAPI(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodAPIRepo.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodAPIDomain, nil).Once()
		result, err := foodService.GetFoodAPI("food test")
		assert.Nil(t, err)
		assert.Equal(t, &foodDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodAPIRepo.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodAPIDomain, assert.AnError).Once()
		result, _ := foodService.GetFoodAPI("food test")
		assert.Equal(t, &foods.Domain{}, result)
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
		mockFoodRepository.On("GetAllFood", mock.Anything, mock.Anything).Return(&[]foods.Domain{}, assert.AnError).Once()
		result, _ := foodService.GetAllFood()
		assert.Equal(t, &[]foods.Domain{}, result)
	})
}

/*	t.Run("Invalid Test", func(t *testing.T) {
	mockFoodRepository.On("GetFoodByName", mock.Anything).Return(&foods.Domain{}, errors.New("hello world")).Once()
	mockFoodRepository.On("Insert", mock.Anything).Return(&foodDomain,nil).Once()
	result, _ := foodService.GetFoodByName("food_test")
	assert.Equal(t, &foodDomain, result)
})*/

// t.Run("Valid Test", func (t *testing.T){})
