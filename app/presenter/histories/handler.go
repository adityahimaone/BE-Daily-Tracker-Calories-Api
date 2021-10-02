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
