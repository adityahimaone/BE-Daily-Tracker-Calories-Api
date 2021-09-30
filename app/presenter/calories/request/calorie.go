package request

import "daily-tracker-calories/bussiness/calories"

type Calorie struct {
	Weight       int    `json:"weight"`
	Height       int    `json:"height"`
	Age          int    `json:"age"`
	ActivityType int    `json:"activity_type"`
	Gender       string `json:"gender"`
}

func ToDomain(request Calorie) *calories.Domain {
	return &calories.Domain{
		Weight:       request.Weight,
		Height:       request.Height,
		Age:          request.Age,
		ActivityType: request.ActivityType,
		Gender:       request.Gender,
	}
}
