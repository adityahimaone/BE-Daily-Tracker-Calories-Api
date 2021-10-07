package histories

import (
	"daily-tracker-calories/bussiness/histories"
	"daily-tracker-calories/repository/mysql/foods"
	"daily-tracker-calories/repository/mysql/users"
	"gorm.io/gorm"
)

type Histories struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	UserID   int
	NameUser string
	Users    users.Users `gorm:"foreignKey:user_id"`
	FoodID   int
	Calorie  float64
	FoodName string
	Foods    foods.Foods `gorm:"foreignKey:food_id"`
	Date     string
}

func toDomain(record Histories) histories.Domain {
	return histories.Domain{
		ID:        int(record.ID),
		UserID:    int(record.Users.ID),
		NameUser:  record.Users.Name,
		FoodID:    int(record.Foods.ID),
		Calorie:   record.Foods.Calorie,
		FoodName:  record.Foods.Name,
		Date:      record.Date,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain histories.Domain) Histories {
	return Histories{
		ID:       uint(domain.ID),
		UserID:   domain.UserID,
		NameUser: domain.NameUser,
		FoodID:   domain.FoodID,
		Calorie:  domain.Calorie,
		FoodName: domain.FoodName,
		Date:     domain.Date,
	}
}

func toDomainArray(record []Histories) []histories.Domain {
	var res []histories.Domain
	for _, v := range record {
		res = append(res, toDomain(v))
	}
	return res
}

type Stat struct {
	SumCalorie float64
}

func toDomainStat(record Stat) histories.Domain {
	return histories.Domain{
		SumCalorie: record.SumCalorie,
	}
}
