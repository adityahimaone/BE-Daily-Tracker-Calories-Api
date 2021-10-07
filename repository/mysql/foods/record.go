package foods

import (
	"daily-tracker-calories/bussiness/foods"
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Name    string
	Calorie float64
	Photo   string
}

func toDomain(record Foods) foods.Domain {
	return foods.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		Calorie:   record.Calorie,
		Photo:     record.Photo,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain foods.Domain) Foods {
	return Foods{
		ID:      uint(domain.ID),
		Name:    domain.Name,
		Calorie: domain.Calorie,
		Photo:   domain.Photo,
	}
}

func toDomainArray(record []Foods) []foods.Domain {
	var res []foods.Domain
	for _, v := range record {
		res = append(res, toDomain(v))
	}
	return res
}
