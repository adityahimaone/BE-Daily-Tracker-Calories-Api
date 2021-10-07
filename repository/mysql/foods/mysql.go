package foods

import (
	"daily-tracker-calories/bussiness/foods"
	"gorm.io/gorm"
)

type repositoryFoods struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) foods.Repository {
	return &repositoryFoods{
		DB: db,
	}
}

func (repository repositoryFoods) GetFoodByID(id int) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Where("id = ?", id).First(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}

func (repository repositoryFoods) GetFoodByName(name string) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Where("name LIKE ?", "%"+name+"%").First(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}

func (repository repositoryFoods) Insert(food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Create(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}

func (repository repositoryFoods) GetAllFood() (*[]foods.Domain, error) {
	var recordFood []Foods
	if err := repository.DB.Find(&recordFood).Error; err != nil {
		return &[]foods.Domain{}, err
	}
	result := toDomainArray(recordFood)
	return &result, nil
}

func (repository repositoryFoods) Delete(id int, food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Where("id = ?", id).Delete(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}

func (repository repositoryFoods) Update(id int, food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Where("id = ?", id).Updates(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}
