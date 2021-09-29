package users

import (
	"daily-tracker-calories/app/middleware/auth"
	_request "daily-tracker-calories/app/presenter/users/request"
	_response "daily-tracker-calories/app/presenter/users/response"
	"daily-tracker-calories/bussiness/users"
	"daily-tracker-calories/helper"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	domain := _request.ToDomainRegister(req)
	resp, err := handler.serviceUser.RegisterUser(domain)
	if err != nil {
		response := helper.APIResponse("Failed Register", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Success Register User", http.StatusOK, "Success", _response.FromDomainRegister(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) LoginUser(echoContext echo.Context) error {
	var req _request.UserLogin
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomainLogin(req)
	resp, err := handler.serviceUser.Login(domain.Email, domain.Password)
	token, err := auth.CreateToken(domain.ID)
	if err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp.Token = token
	fmt.Println(resp.ID)
	if err != nil {
		response := helper.APIResponse("Generate Token Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success Login", http.StatusOK, "Success", _response.FromDomainLogin(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceUser.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomainRegister(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
