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
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *serviceUsers) Update(user *Domain) (*Domain, error) {
	panic("implement me")
}

func (s *serviceUsers) FindByID(id int) (*Domain, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return user, nil
}

func (s *serviceUsers) Login(email string, password string) (*Domain, error) {
	result, err := s.repository.Login(email, password)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
