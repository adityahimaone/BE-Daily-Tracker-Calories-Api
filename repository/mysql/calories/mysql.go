package calories

import (
	"daily-tracker-calories/bussiness/calories"
	"gorm.io/gorm"
	"log"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) calories.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Insert(calorie *calories.Domain, idUser int) (*calories.Domain, error) {
	recordCalorie := fromDomain(*calorie)
	log.Println(idUser)
	if err := repository.DB.Where("user_id = ?", idUser).Create(&recordCalorie).Error; err != nil {
		return &calories.Domain{}, err
	}
	result := toDomain(recordCalorie)
	return &result, nil
}

func (repository repositoryUsers) Update(calorie *calories.Domain, id int) (*calories.Domain, error) {
	recordCalorie := fromDomain(*calorie)
	log.Println(id)
	if err := repository.DB.Where("user_id = ?", id).Updates(&recordCalorie).Error; err != nil {
		return &calories.Domain{}, err
	}
	result := toDomain(recordCalorie)
	return &result, nil
}

func (repository repositoryUsers) GetCalorieByUserID(id int) (*calories.Domain, error) {
	recordCalorie := Calories{}
	if err := repository.DB.Where("user_id = ?", id).First(&recordCalorie).Error; err != nil {
		return &calories.Domain{}, err
	}
	result := toDomain(recordCalorie)
	return &result, nil
}
