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

func (service *serviceUsers) RegisterUser(user *Domain) (*Domain, error) {
	passwordHash, err := helper.PasswordHash(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = passwordHash
	_, err = service.repository.FindByEmail(user.Email)
	if err != nil {
		result, err := service.repository.Insert(user)
		if err != nil {
			return &Domain{}, err
		}
		return result, nil
	}
	return &Domain{}, errors.New("email registered")
}

func (service *serviceUsers) Update(id int, user *Domain) (*Domain, error) {
	passwordHash, err := helper.PasswordHash(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = passwordHash
	result, err := service.repository.Update(id, user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceUsers) FindByID(id int) (*Domain, error) {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return user, nil
}

func (service *serviceUsers) Login(email string, password string) (string, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return "ID Not Found", errors.New("User Not Found")
	}
	if user.ID == 0 {
		return "ID Not Found", errors.New("User Not Found")
	}
	if !helper.ValidateHash(password, user.Password) {
		return "Error Validate Hash", errors.New("Error Validate Hash")
	}
	token := service.jwtAuth.GenerateToken(user.ID)
	return token, nil
}

func (service *serviceUsers) UploadAvatar(id int, fileLocation string) (*Domain, error) {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	user.Avatar = fileLocation
	updateAvatar, err := service.repository.Update(id, user)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (service *serviceUsers) EmailAvailable(email string) (bool, error) {
	user, _ := service.repository.FindByEmail(email)
	if user != nil {
		return false, nil
	}
	return true, nil
}
