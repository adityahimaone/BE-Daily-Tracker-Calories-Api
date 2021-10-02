package response

import (
	"daily-tracker-calories/bussiness/histories"
	"time"
)

type Histories struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	NameUser  string    `json:"name_user"`
	FoodID    int       `json:"food_id"`
	Calorie   float64   `json:"calorie"`
	FoodName  string    `json:"food_name"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain histories.Domain) Histories {
	return Histories{
		ID:        domain.ID,
		UserID:    domain.UserID,
		NameUser:  domain.NameUser,
		FoodID:    domain.FoodID,
		Calorie:   domain.Calorie,
		FoodName:  domain.FoodName,
		Date:      domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
