package histories

import (
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	NameUser  string
	FoodID    int
	Calorie   float64
	FoodName  string
	Date      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateHistories(histories *Domain) (*Domain, error)
	GetAllHistoriesByUserID(userid int) (*[]Domain, error)
}

type Repository interface {
	Insert(histories *Domain) (*Domain, error)
	GetHistoryByUserID(userid int) (*Domain, error)
	GetAllHistoriesByUserID(userid int) (*[]Domain, error)
}
