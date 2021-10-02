package histories

import (
	"daily-tracker-calories/bussiness/foods"
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	NameUser  string
	FoodID    int
	Calorie   float64
	FoodName  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateHistories(histories *Domain) (*Domain, error)
	GetHistoriesByUserID(userid int) (*[]Domain, error)
}

type Repository interface {
	Insert(histories *Domain, foods *foods.Domain) (*Domain, error)
}
