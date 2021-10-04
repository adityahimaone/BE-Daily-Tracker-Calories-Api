package foods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type FoodNutrients struct {
	Title  string  `json:"title"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}
type FoodNutrion struct {
	FoodNutrients FoodNutrients
}
type FoodDesc struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	ImageType   string `json:"imageType"`
	FoodNutrion []FoodNutrion
}
type Food struct {
	Results      []FoodDesc `json:"results"`
	Offset       int        `json:"offset"`
	Number       int        `json:"number"`
	TotalResults int        `json:"totalResults"`
}

type AutoGenerated struct {
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

func TestGetFood(t *testing.T) {
	apikey := "f85868cf3e9f448c851d46fe687a40ac"
	query := "cup cake"
	splitQuery := strings.Split(query, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	log.Println(joinQuery)
	minCal := 0
	number := 1
	urlString := fmt.Sprintf("https://api.foodAPI.com/recipes/complexSearch?apiKey=%s&query=%s&minCalories=%d&number=%d", apikey, joinQuery, minCal, number)
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
	food := AutoGenerated{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}
	var array1 []struct {
		Title  string
		Image  string
		Amount float64
	}
	for _, valueResult := range food.Results {
		for _, valueNutrients := range valueResult.Nutrition.Nutrients {
			array1 = append(array1, struct {
				Title  string
				Image  string
				Amount float64
			}{Title: valueResult.Title, Image: valueResult.Image, Amount: valueNutrients.Amount})
			fmt.Println(array1)
		}
	}
	titleFood := ""
	imageFood := ""
	amountKcal := 0.0
	var valueFood []string
	var valueCal []float64
	for _, v := range array1 {
		valueFood = append(valueFood, v.Title, v.Image)
	}
	for _, v := range array1 {
		valueCal = append(valueCal, v.Amount)
	}
	titleFood = valueFood[0]
	imageFood = valueFood[1]
	amountKcal = valueCal[0]
	fmt.Println(titleFood)
	fmt.Println(imageFood)
	fmt.Println(amountKcal)
}

func TestSplit(t *testing.T) {
	query := "cup cake"
	splitQuery := strings.Split(query, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	log.Println(joinQuery)

	foodNts := FoodNutrients{"Calories", "calories", 200.0, "kcal"}
	foodNto := FoodNutrion{foodNts}
	food1 := FoodDesc{12, "Anggur", "", "", []FoodNutrion{foodNto}}
	res := Food{[]FoodDesc{food1}, 1, 2, 3}
	log.Println(res)
}

/*var array1 []struct
for keyResult, valueResult := range struct.Results{
for keyNutrients, valueNutiens := range valuevalueResult.Nutrition.Nutriciens{
array1 := append(array1, struct{
Title : valueResult.Title
Image :
NutitionTitle : valueNutiens.Title
})
}
}
*/
