package users

import (
	"daily-tracker-calories/app/middleware/auth"
	"daily-tracker-calories/helper"
	"errors"
)

type serviceUsers struct {
	repository Repository
	jwtAuth    *auth.ConfigJWT
}

func NewService(repositoryUser Repository, jwtauth *auth.ConfigJWT) Service {
	return &serviceUsers{
		repository: repositoryUser,
		jwtAuth:    jwtauth,
	}
}

func (s *serviceUsers) RegisterUser(user *Domain) (*Domain, error) {
	passwordHash, err := helper.PasswordHash(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = passwordHash
	result, err := s.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceUsers) Update(id int, user *Domain) (*Domain, error) {
	passwordHash, err := helper.PasswordHash(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = passwordHash
	result, err := s.repository.Update(id, user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceUsers) FindByID(id int) (*Domain, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return user, nil
}

func (s *serviceUsers) Login(email string, password string) (string, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return "ID Not Found", errors.New("User Not Found")
	}
	if user.ID == 0 {
		return "ID Not Found", errors.New("User Not Found")
	}
	if !helper.ValidateHash(password, user.Password) {
		return "Error Validate Hash", errors.New("Error Validate Hash")
	}
	token := s.jwtAuth.GenerateToken(user.ID)
	return token, nil
}

func (s *serviceUsers) UploudAvatar(id int, fileLocation string) (*Domain, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	user.Avatar = fileLocation
	updateAvatar, err := s.repository.Update(id, user)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}
