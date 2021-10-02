package calories

import (
	"daily-tracker-calories/app/middleware/auth"
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

func (handler *Presenter) SaveCalorie(echoContext echo.Context) error {
	req := _request.Calorie{}
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	user := auth.GetUser(echoContext) // ID Get From JWT
	userID := user.ID
	resp, err := handler.serviceCalorie.CreateCalorie(domain, userID)
	if err != nil {
		response := helper.APIResponse("Failed Get Calorie", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get Calorie", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
