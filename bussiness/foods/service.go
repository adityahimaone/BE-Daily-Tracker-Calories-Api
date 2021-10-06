package foods

import (
	"daily-tracker-calories/bussiness/foodAPI"
)

type serviceFoods struct {
	repository  Repository
	foodAPIRepo foodAPI.Repository
}

func NewService(repositoryFood Repository, foodAPIRepo foodAPI.Repository) Service {
	return &serviceFoods{
		repository:  repositoryFood,
		foodAPIRepo: foodAPIRepo,
	}
}

func (service *serviceFoods) GetFoodByID(id int) (*Domain, error) {
	result, err := service.repository.GetFoodByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceFoods) GetFoodByName(name string) (*Domain, error) {
	result, err := service.repository.GetFoodByName(name)
	if err != nil {
		apiFood, err := service.GetFoodAPI(name)
		if err != nil {
			return &Domain{}, err
		}
		insert, err := service.SaveFood(apiFood)
		if err != nil {
			return &Domain{}, err
		}
		return insert, nil
	}
	return result, nil
}

func (service *serviceFoods) GetFoodAPI(name string) (*Domain, error) {
	result, err := service.foodAPIRepo.GetFoodByName(name)
	if err != nil {
		return &Domain{}, err
	}
	newRes := Domain{
		Name:    result.Name,
		Calorie: result.Calorie,
		Photo:   result.Photo,
	}
	return &newRes, nil
}

func (service *serviceFoods) SaveFood(food *Domain) (*Domain, error) {
	result, err := service.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceFoods) GetAllFood() (*[]Domain, error) {
	result, err := service.repository.GetAllFood()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}

func (service *serviceFoods) DeleteFood(id int, food *Domain) (*Domain, error) {
	result, err := service.repository.Delete(id, food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceFoods) EditFood(id int, food *Domain) (*Domain, error) {
	result, err := service.repository.Update(id, food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
