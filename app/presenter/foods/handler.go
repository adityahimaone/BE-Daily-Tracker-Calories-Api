package foods

import (
	_request "daily-tracker-calories/app/presenter/foods/request"
	_response "daily-tracker-calories/app/presenter/foods/response"
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	response := helper.APIResponse("Success Save Food", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) GetFoodByName(echoContext echo.Context) error {
	name := echoContext.QueryParam("name")
	resp, err := handler.serviceFood.GetFoodByName(name)
	if err != nil {
		response := helper.APIResponse("Failed Get Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) GetAllFood(echoContext echo.Context) error {
	resp, err := handler.serviceFood.GetAllFood()
	if err != nil {
		response := helper.APIResponse("Failed Get All Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomainArray(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) GetFoodByID(echoContext echo.Context) error {
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helper.APIResponse("Failed Get Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceFood.GetFoodByID(id)
	if err != nil {
		response := helper.APIResponse("Failed Get Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) DeleteFood(echoContext echo.Context) error {
	var req _request.Food
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Delete Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helper.APIResponse("Failed Delete Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	_, err = handler.serviceFood.DeleteFood(id, domain)
	if err != nil {
		response := helper.APIResponse("Failed Delete Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.Delete{Data: "Success Delete Food"})
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) EditFood(echoContext echo.Context) error {
	var req _request.Food
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Edit Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helper.APIResponse("Failed Edit Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceFood.EditFood(id, domain)
	resp.ID = id
	if err != nil {
		response := helper.APIResponse("Failed Edit Food", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)

}
