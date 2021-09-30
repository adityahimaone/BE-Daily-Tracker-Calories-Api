package routes

import (
	"daily-tracker-calories/app/presenter/calories"
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserHandler    users.Presenter
	JWTMiddleware  middleware.JWTConfig
	CalorieHandler calories.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")
	//group.Use(middleware.JWT([]byte(viper.GetString(`jwt.token`))))
	group.POST("/user/register", handler.UserHandler.RegisterUser)
	group.PUT("/user/update", handler.UserHandler.UpdateUser, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.POST("/user/login", handler.UserHandler.LoginUser)
	group.GET("/user/:id", handler.UserHandler.FindByID)

	//calorie endpoint
	group.POST("/calorie/count", handler.CalorieHandler.CountCalorie)

}
