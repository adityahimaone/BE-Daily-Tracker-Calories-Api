package foods

type serviceFoods struct {
	repository Repository
}

func NewService(repositoryFood Repository) Service {
	return &serviceFoods{
		repository: repositoryFood,
	}
}

func (s serviceFoods) GetFoodByName(name string) (*Domain, error) {
	panic("implement me")
}

func (s serviceFoods) GetFoodAPI(food *Domain) (*Domain, error) {
	panic("implement me")
}

func (s serviceFoods) SaveFood(food *Domain) (*Domain, error) {
	result, err := s.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
