package spoonacular

import (
	"daily-tracker-calories/bussiness/foodAPI"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type spoonacularAPI struct {
	httpClient http.Client
}

func NewFoodAPI() foodAPI.Repository {
	return &spoonacularAPI{
		httpClient: http.Client{},
	}
}

func (s spoonacularAPI) GetFoodByName(name string) (*foodAPI.Domain, error) {
	apikey := "f85868cf3e9f448c851d46fe687a40ac"
	//query := "cup cake"
	splitQuery := strings.Split(name, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	log.Println(joinQuery)
	minCal := 0
	number := 1
	urlString := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&minCalories=%d&number=%d", apikey, joinQuery, minCal, number)
	log.Println(urlString)
	response, err := http.Get(urlString)
	if err != nil {
		log.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	food := FoodsSource{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}
	var foodTemp []struct {
		Title  string
		Image  string
		Amount float64
	}
	for _, valueResult := range food.Results {
		for _, valueNutrients := range valueResult.Nutrition.Nutrients {
			foodTemp = append(foodTemp, struct {
				Title  string
				Image  string
				Amount float64
			}{Title: valueResult.Title, Image: valueResult.Image, Amount: valueNutrients.Amount})
			fmt.Println(foodTemp)
		}
	}
	titleFood := ""
	imageFood := ""
	amountKcal := 0.0
	var valueFood []string
	var valueCal []float64
	for _, v := range foodTemp {
		valueFood = append(valueFood, v.Title, v.Image)
	}
	for _, v := range foodTemp {
		valueCal = append(valueCal, v.Amount)
	}
	titleFood = valueFood[0]
	imageFood = valueFood[1]
	amountKcal = valueCal[0]
	result := toDomain(Foods{Title: titleFood, Image: imageFood, Amount: amountKcal})
	return &result, nil
}
