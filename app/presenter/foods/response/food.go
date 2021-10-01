package response

import (
	"daily-tracker-calories/bussiness/foods"
	"time"
)

type Food struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Calorie   float64   `json:"calorie"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain foods.Domain) Food {
	return Food{
		ID:        domain.ID,
		Name:      domain.Name,
		Calorie:   domain.Calorie,
		Photo:     domain.Photo,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
