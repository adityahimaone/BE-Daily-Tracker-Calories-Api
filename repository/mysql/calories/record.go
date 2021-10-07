package calories

import (
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/repository/mysql/users"
	"gorm.io/gorm"
)

type Calories struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Calorie float64
	UserID  int
	Users   users.Users `gorm:"foreignKey:user_id"`
}

func toDomain(record Calories) calories.Domain {
	return calories.Domain{
		ID:      int(record.ID),
		Calorie: record.Calorie,
		UserID:  record.UserID,
	}
}

func fromDomain(domain calories.Domain) Calories {
	return Calories{
		ID:      uint(domain.ID),
		Calorie: domain.Calorie,
		UserID:  domain.UserID,
	}
}
