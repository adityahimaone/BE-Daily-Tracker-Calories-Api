package routes

import (
	"daily-tracker-calories/app/presenter/calories"
	"daily-tracker-calories/app/presenter/foods"
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserHandler    users.Presenter
	JWTMiddleware  middleware.JWTConfig
	CalorieHandler calories.Presenter
	FoodHandler    foods.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")
	//group.Use(middleware.JWT([]byte(viper.GetString(`jwt.token`))))
	group.POST("/users/register", handler.UserHandler.RegisterUser)
	group.PUT("/users/update", handler.UserHandler.UpdateUser, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.POST("/users/login", handler.UserHandler.LoginUser)
	group.GET("/users/:id", handler.UserHandler.FindByID)

	//calorie endpoint
	group.POST("/calorie/count", handler.CalorieHandler.CountCalorie)
	group.POST("/calorie/save", handler.CalorieHandler.SaveCalorie, middleware.JWTWithConfig(handler.JWTMiddleware))

	//food endpoint
	group.POST("/food/save", handler.FoodHandler.SaveFood)
}
