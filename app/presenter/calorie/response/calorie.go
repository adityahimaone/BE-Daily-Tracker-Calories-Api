package response

import "daily-tracker-calories/bussiness/calories"

type Calorie struct {
	ID            int `json:"id"`
	ResultCalorie int `json:"result_calorie"`
}

func FromDomain(domain calories.Domain) Calorie {
	return Calorie{
		ID:            domain.ID,
		ResultCalorie: domain.Calorie,
	}
}
