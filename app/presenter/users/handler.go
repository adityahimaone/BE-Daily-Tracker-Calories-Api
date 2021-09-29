package users

import (
	"daily-tracker-calories/app/middleware/auth"
	_request "daily-tracker-calories/app/presenter/users/request"
	_response "daily-tracker-calories/app/presenter/users/response"
	"daily-tracker-calories/bussiness/users"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"log"
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
	var req _request.User
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Register", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceUser.RegisterUser(domain)
	if err != nil {
		response := helper.APIResponse("Failed Register", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("Success Register User", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) LoginUser(echoContext echo.Context) error {
	req := _request.UserLogin{}
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceUser.Login(req.Email, req.Password)
	token, err := auth.CreateToken(resp.ID)
	resp.Token = token
	if err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := helper.APIResponse("Generate Token Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success Login", http.StatusOK, "Success", _response.UserLogin{ID: resp.ID, Email: resp.Email, Token: resp.Token})
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) UpdateUser(echoContext echo.Context) error {
	var req _request.User
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	user := auth.GetUser(echoContext)
	log.Println(user.ID)
	resp, err := handler.serviceUser.Update(user.ID, domain)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceUser.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
