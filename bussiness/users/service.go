package users

import (
	"daily-tracker-calories/helper"
	"errors"
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

func (s *serviceUsers) Login(email string, password string) (*Domain, error) {
	user, err := s.repository.Login(email, password)
	if err != nil {
		return &Domain{}, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}
	if !helper.ValidateHash(password, user.Password) {
		return user, err
	}
	log.Println(user)
	return user, nil
}
