package calories

import (
	"daily-tracker-calories/app/middleware/auth"
	"daily-tracker-calories/app/middleware/validate"
	_request "daily-tracker-calories/app/presenter/calories/request"
	_response "daily-tracker-calories/app/presenter/calories/response"
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Presenter struct {
	serviceCalorie calories.Service
	validate       *validator.Validate
}

func NewHandler(calorieServ calories.Service, validator *validator.Validate) *Presenter {
	return &Presenter{
		serviceCalorie: calorieServ,
		validate:       validator,
	}
}

func (handler *Presenter) CountCalorie(echoContext echo.Context) error {
	req := _request.Calorie{}
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Get Calorie", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	if err := handler.validate.Struct(req); err != nil {
		errors := validate.FormatValidationError(err)
		errorsData := map[string]interface{}{"errors": errors}
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error Validation", errorsData)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
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
	if err := handler.validate.Struct(req); err != nil {
		errors := validate.FormatValidationError(err)
		errorsData := map[string]interface{}{"errors": errors}
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error Validation", errorsData)
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

func (handler *Presenter) GetCalorieByUserID(echoContext echo.Context) error {
	user := auth.GetUser(echoContext) // ID Get From JWT
	userID := user.ID
	resp, err := handler.serviceCalorie.GetCalorieByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed Get Calorie By User ID", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get Calorie", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
