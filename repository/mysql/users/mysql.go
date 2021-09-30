package users

import (
	"daily-tracker-calories/bussiness/users"
	"gorm.io/gorm"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) Update(id int, user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Where("id = ?", id).Updates(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) FindByID(id int) (*users.Domain, error) {
	var recordUser Users
	if err := repository.DB.Where("id = ?", id).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) Login(email string, password string) (*users.Domain, error) {
	recordUser := Users{}
	if err := repository.DB.Where("email = ? AND password = ?", email, password).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}
