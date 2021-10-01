package foods

import (
	_request "daily-tracker-calories/app/presenter/foods/request"
	_response "daily-tracker-calories/app/presenter/foods/response"
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Presenter struct {
	serviceFood foods.Service
}

func NewHandler(foodService foods.Service) *Presenter {
	return &Presenter{
		serviceFood: foodService,
	}
}

func (handler *Presenter) SaveFood(echoContext echo.Context) error {
	var req _request.Food
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Save Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceFood.SaveFood(domain)
	if err != nil {
		response := helper.APIResponse("Failed Save Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Register User", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}
