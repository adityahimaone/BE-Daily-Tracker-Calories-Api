package users

type serviceUsers struct {
	repository Repository
}

func NewService(repositoryUser Repository) Service {
	return &serviceUsers{
		repository: repositoryUser,
	}
}

func (s *serviceUsers) RegisterUser(user *Domain) (*Domain, error) {
	result, err := s.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceUsers) IsEmailAvailable(email string) (bool, error) {
	panic("implement me")
}

func (s *serviceUsers) Update(user *Domain) (*Domain, error) {
	panic("implement me")
}

func (s *serviceUsers) FindByID(id int) (*Domain, error) {
	panic("implement me")
}

func (s *serviceUsers) Login(username string, password string) (*Domain, error) {
	panic("implement me")
}
