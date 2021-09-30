package calories

import (
	"daily-tracker-calories/bussiness/calories"
	"gorm.io/gorm"
)

type Calories struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Calorie float64
}

func toDomain(record Calories) calories.Domain {
	return calories.Domain{
		ID:      int(record.ID),
		Calorie: record.Calorie,
	}
}

func fromDomain(domain calories.Domain) Calories {
	return Calories{
		ID:      uint(domain.ID),
		Calorie: domain.Calorie,
	}
}
