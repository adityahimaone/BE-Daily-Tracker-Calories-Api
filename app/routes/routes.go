package routes

import (
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserHandler   users.Presenter
	JWTMiddleware middleware.JWTConfig
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	users := e.Group("/api/v1")
	//users.Use(middleware.JWT([]byte(viper.GetString(`jwt.token`))))
	users.POST("/users/register", handler.UserHandler.RegisterUser)
	users.PUT("/users/update", handler.UserHandler.UpdateUser, middleware.JWTWithConfig(handler.JWTMiddleware))
	users.POST("/users/login", handler.UserHandler.LoginUser)
	users.GET("/users/:id", handler.UserHandler.FindByID)
}
