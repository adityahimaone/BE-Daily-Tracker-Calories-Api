package routes

import (
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	UserHandler users.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", handler.UserHandler.RegisterUser)
}
