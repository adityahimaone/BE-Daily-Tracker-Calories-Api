package request

import "daily-tracker-calories/bussiness/calories"

type Calorie struct {
	Weight       int    `json:"weight" validate:"required,gte=0"`
	Height       int    `json:"height" validate:"required,gte=0"`
	Age          int    `json:"age" validate:"required,gte=0,lte=150"`
	ActivityType int    `json:"activity_type" validate:"required,gte=1,lte=5"`
	Gender       string `json:"gender" validate:"required"`
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
