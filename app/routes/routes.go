package routes

import (
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	UserHandler users.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	users := e.Group("/api/v1")
	users.POST("/users/register", handler.UserHandler.RegisterUser)
	users.POST("/users/login", handler.UserHandler.LoginUser)
	users.GET("/users/:id", handler.UserHandler.FindByID)
}
