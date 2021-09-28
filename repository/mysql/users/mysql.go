package users

import (
	"daily-tracker-calories/bussiness/users"
	"gorm.io/gorm"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	return user, nil
}

func (repository repositoryUsers) Update(user *users.Domain) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) FindByID(id int) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) FindByEmail(email string) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) Login(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Where("email = ? & password ?", recordUser.Email, recordUser.Password).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}
