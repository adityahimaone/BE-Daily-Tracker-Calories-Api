package foods

type serviceFoods struct {
	repository Repository
}

func NewService(repositoryFood Repository) Service {
	return &serviceFoods{
		repository: repositoryFood,
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
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceFoods) GetFoodAPI(food *Domain) (*Domain, error) {
	panic("implement me")
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
