package users

import (
	"daily-tracker-calories/app/middleware/auth"
	_request "daily-tracker-calories/app/presenter/users/request"
	_response "daily-tracker-calories/app/presenter/users/response"
	"daily-tracker-calories/bussiness/users"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Presenter struct {
	serviceUser users.Service
	authService auth.Service
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
	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	req.Password = string(passwordHash)
	domain := _request.ToDomainRegister(req)
	resp, err := handler.serviceUser.RegisterUser(domain)
	if err != nil {
		response := helper.APIResponse("Failed Register", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
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

	resp, err := handler.serviceUser.Login(domain)
	//token, err := handler.authService.GenerateToken(resp.ID)
	//resp.Token = token
	if err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success Login", http.StatusOK, "Success", _response.FromDomainLogin(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
