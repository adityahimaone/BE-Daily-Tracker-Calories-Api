package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type serviceUsers struct {
	repository Repository
}

func NewService(repositoryUser Repository) Service {
	return &serviceUsers{
		repository: repositoryUser,
	}
}

func (s *serviceUsers) RegisterUser(user *Domain) (*Domain, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(passwordHash)
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
	user, err := s.repository.Login(email, password)
	if err != nil {
		return &Domain{}, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	log.Println(user)
	return user, nil
}
