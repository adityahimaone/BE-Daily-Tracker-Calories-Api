package users

import (
	_request "daily-tracker-calories/app/presenter/users/request"
	_response "daily-tracker-calories/app/presenter/users/response"
	"daily-tracker-calories/bussiness/users"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Presenter struct {
	serviceUser users.Service
}

func NewHandler(userService users.Service) *Presenter {
	return &Presenter{
		serviceUser: userService,
	}
}

func (handler *Presenter) RegisterUser(echoContext echo.Context) error {
	var req _request.UserRegister
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Register", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	domain := _request.ToDomain(req)
	resp, err := handler.serviceUser.RegisterUser(domain)
	if err != nil {
		response := helper.APIResponse("Failed Register", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success Register User", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
