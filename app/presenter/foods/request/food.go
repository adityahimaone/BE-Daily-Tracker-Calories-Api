package request

import "daily-tracker-calories/bussiness/foods"

type Food struct {
	Name    string  `json:"name"`
	Calorie float64 `json:"calorie"`
	Photo   string  `json:"photo"`
}

func ToDomain(request Food) *foods.Domain {
	return &foods.Domain{
		Name:    request.Name,
		Calorie: request.Calorie,
		Photo:   request.Photo,
	}
}
