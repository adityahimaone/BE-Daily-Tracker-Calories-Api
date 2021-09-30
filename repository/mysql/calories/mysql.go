package calories

import (
	"daily-tracker-calories/bussiness/calories"
	"gorm.io/gorm"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) calories.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (r repositoryUsers) Insert(calorie *calories.Domain) (*calories.Domain, error) {
	panic("implement me")
}

func (r repositoryUsers) Update(calorie *calories.Domain) (*calories.Domain, error) {
	panic("implement me")
}
