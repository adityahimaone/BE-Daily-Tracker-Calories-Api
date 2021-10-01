package response

import "daily-tracker-calories/bussiness/calories"

type Calorie struct {
	ID            int     `json:"id"`
	ResultCalorie float64 `json:"result_calorie"`
	UserID        int     `json:"user_id"`
}

type CalorieResult struct {
	ResultCalorie float64 `json:"result_calorie"`
}

func FromDomain(domain calories.Domain) Calorie {
	return Calorie{
		ID:            domain.ID,
		ResultCalorie: domain.Calorie,
		UserID:        domain.UserID,
	}
}
