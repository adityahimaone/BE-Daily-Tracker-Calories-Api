package histories

import (
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

func (repository repositoryHistories) Insert(history *histories.Domain) (*histories.Domain, error) {
	recordHistory := fromDomain(*history)
	if err := repository.DB.Create(&recordHistory).Error; err != nil {
		return &histories.Domain{}, err
	}
	result, err := repository.GetHistoryByUserID(recordHistory.UserID)
	if err != nil {
		return &histories.Domain{}, err
	}
	return result, nil
}

func (repository repositoryHistories) GetHistoryByUserID(userid int) (*histories.Domain, error) {
	recordHistory := Histories{}
	if err := repository.DB.Where("user_id = ?", userid).Joins("Users").Joins("Foods").Last(&recordHistory).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := toDomain(recordHistory)
	return &result, nil
}

func (repository repositoryHistories) GetAllHistoriesByUserID(userid int) (*[]histories.Domain, error) {
	var recordHistory []Histories
	if err := repository.DB.Where("user_id = ?", userid).Joins("Users").Joins("Foods").Find(&recordHistory).Group("date").Error; err != nil {
		return &[]histories.Domain{}, err
	}
	result := toDomainArray(recordHistory)
	return &result, nil
}
