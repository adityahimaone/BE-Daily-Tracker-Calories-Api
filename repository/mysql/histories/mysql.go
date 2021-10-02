package histories

import (
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/bussiness/histories"
	"gorm.io/gorm"
)

type repositoryHistories struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) histories.Repository {
	return &repositoryHistories{
		DB: db,
	}
}

func (r repositoryHistories) Insert(history *histories.Domain, foods *foods.Domain) (*histories.Domain, error) {
	recordHistory := fromDomain(*history)
	if err := r.DB.Create(&recordHistory).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := toDomain(recordHistory)
	return &result, nil
}
