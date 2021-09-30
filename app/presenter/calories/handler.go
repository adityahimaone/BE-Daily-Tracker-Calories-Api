package calories

import (
	_request "daily-tracker-calories/app/presenter/calories/request"
	_response "daily-tracker-calories/app/presenter/calories/response"
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Presenter struct {
	serviceCalorie calories.Service
}

func NewHandler(calorieServ calories.Service) *Presenter {
	return &Presenter{
		serviceCalorie: calorieServ,
	}
}

func (handler *Presenter) CountCalorie(echoContext echo.Context) error {
	req := _request.Calorie{}
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Get Calorie", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	log.Println(domain)
	_, err := handler.serviceCalorie.CountCalorie(domain)
	if err != nil {
		response := helper.APIResponse("Failed Get Calorie", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get Calorie", http.StatusOK, "Success", _response.CalorieResult{ResultCalorie: domain.Calorie})
	return echoContext.JSON(http.StatusOK, response)
}
