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

func (s *serviceFoods) GetFoodByID(id int) (*Domain, error) {
	result, err := s.repository.GetFoodByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceFoods) GetFoodByName(name string) (*Domain, error) {
	result, err := s.repository.GetFoodByName(name)
	if err != nil {
		apiFood, err := s.GetFoodAPI(name)
		if err != nil {
			return &Domain{}, err
		}
		insert, err := s.SaveFood(apiFood)
		if err != nil {
			return &Domain{}, err
		}
		return insert, nil
	}
	return result, nil
}

func (s *serviceFoods) GetFoodAPI(name string) (*Domain, error) {
	result, err := s.foodAPIRepo.GetFoodByName(name)
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

func (s *serviceFoods) SaveFood(food *Domain) (*Domain, error) {
	result, err := s.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceFoods) GetAllFood() (*[]Domain, error) {
	result, err := s.repository.GetAllFood()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}

func (s *serviceFoods) DeleteFood(id int, food *Domain) (*Domain, error) {
	result, err := s.repository.Delete(id, food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

/*func (s *serviceFoods) GetFoodByName(name string) (*Domain, error) {
	result, err := s.repository.GetFoodByName(name)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}*/

/*func (s *serviceFoods) SaveFood(food *Domain) (*Domain, error) {
	result, err := s.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}*/
