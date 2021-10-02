package request

import "daily-tracker-calories/bussiness/histories"

type Histories struct {
	FoodName string `json:"food_name"`
}

func ToDomain(request Histories) *histories.Domain {
	return &histories.Domain{
		FoodName: request.FoodName,
	}
}
