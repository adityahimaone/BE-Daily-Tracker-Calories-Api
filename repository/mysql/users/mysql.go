package users

import (
	"daily-tracker-calories/bussiness/users"
	"gorm.io/gorm"
)

type repositoryUsers struct{
	DB *gorm.DB
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Create(&recordUser).Error; err != nil{
		return &users.Domain{}, err
	}
	return user, nil
}

func (r repositoryUsers) Update(user *users.Domain) (*users.Domain, error) {
	panic("implement me")
}

func (r repositoryUsers) FindByID(id int) (*users.Domain, error) {
	panic("implement me")
}

func (r repositoryUsers) FindByEmail(email string) (*users.Domain, error) {
	panic("implement me")
}

func (r repositoryUsers) Login(username string, password string) (*users.Domain, error) {
	panic("implement me")
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository{
	return &repositoryUsers{
		DB: db,
	}
}