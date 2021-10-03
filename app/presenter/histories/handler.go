package histories

import (
	"daily-tracker-calories/app/middleware/auth"
	_request "daily-tracker-calories/app/presenter/histories/request"
	_response "daily-tracker-calories/app/presenter/histories/response"
	"daily-tracker-calories/bussiness/histories"
	"daily-tracker-calories/helper"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Presenter struct {
	serviceHistories histories.Service
}

func NewHandler(historiesService histories.Service) *Presenter {
	return &Presenter{
		serviceHistories: historiesService,
	}
}

func (handler *Presenter) CreateHistory(echoContext echo.Context) error {
	var req _request.Histories
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	user := auth.GetUser(echoContext)
	domain.UserID = user.ID
	log.Println("domain", domain)
	resp, err := handler.serviceHistories.CreateHistories(domain)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get Calorie", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) GetAllHistoriesByUserID(echoContext echo.Context) error {
	user := auth.GetUser(echoContext)
	userID := user.ID
	resp, err := handler.serviceHistories.GetAllHistoriesByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get All Food By User ID", http.StatusBadRequest, "Success", _response.FromDomainArray(*resp))
	return echoContext.JSON(http.StatusBadRequest, response)
}

func (handler *Presenter) UserStat(echoContext echo.Context) error {
	user := auth.GetUser(echoContext)
	userID := user.ID
	current, need, status, err := handler.serviceHistories.UserStat(userID)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Get Calorie", http.StatusOK, "Success", _response.UserStat{CalorieCurrent: current, CalorieNeed: need, Status: status})
	return echoContext.JSON(http.StatusOK, response)
}
