package spoonacular

import "daily-tracker-calories/bussiness/foodAPI"

type FoodsSource struct {
	Results []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Image     string `json:"image"`
		ImageType string `json:"imageType"`
		Nutrition struct {
			Nutrients []struct {
				Title  string  `json:"title"`
				Name   string  `json:"name"`
				Amount float64 `json:"amount"`
				Unit   string  `json:"unit"`
			} `json:"nutrients"`
		} `json:"nutrition"`
	} `json:"results"`
	Offset       int `json:"offset"`
	Number       int `json:"number"`
	TotalResults int `json:"totalResults"`
}

type Foods struct {
	Title  string
	Image  string
	Amount float64
}

func toDomain(record Foods) foodAPI.Domain {
	return foodAPI.Domain{
		Name:    record.Title,
		Photo:   record.Image,
		Calorie: record.Amount,
	}
}
