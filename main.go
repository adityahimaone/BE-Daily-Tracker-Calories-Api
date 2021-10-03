package main

import (
	_middleware "daily-tracker-calories/app/middleware"
	"daily-tracker-calories/app/middleware/auth"
	_handlerCalories "daily-tracker-calories/app/presenter/calories"
	_handlerFoods "daily-tracker-calories/app/presenter/foods"
	_handlerHistories "daily-tracker-calories/app/presenter/histories"
	_handlerUsers "daily-tracker-calories/app/presenter/users"
	"daily-tracker-calories/app/routes"
	_serviceCalories "daily-tracker-calories/bussiness/calories"
	_serviceFoods "daily-tracker-calories/bussiness/foods"
	_serviceHistories "daily-tracker-calories/bussiness/histories"
	_serviceUsers "daily-tracker-calories/bussiness/users"
	mysqlRepo "daily-tracker-calories/repository/mysql"
	_repositoryCalories "daily-tracker-calories/repository/mysql/calories"
	_repositoryFoods "daily-tracker-calories/repository/mysql/foods"
	_repositoryHistories "daily-tracker-calories/repository/mysql/histories"
	_repositoryUsers "daily-tracker-calories/repository/mysql/users"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := mysqlRepo.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	db := configDB.IntialDB()
	mysqlRepo.MigrateDB(db)

	e := echo.New()

	//factory of domain
	//authService := _middleware.NewHandler()
	userRepository := _repositoryUsers.NewRepositoryMySQL(db)
	userService := _serviceUsers.NewService(userRepository, &configJWT)
	usersHandler := _handlerUsers.NewHandler(userService, &configJWT)

	calorieRepository := _repositoryCalories.NewRepositoryMySQL(db)
	calorieService := _serviceCalories.NewService(calorieRepository, userService)
	calorieHandler := _handlerCalories.NewHandler(calorieService)

	foodRepository := _repositoryFoods.NewRepositoryMySQL(db)
	foodService := _serviceFoods.NewService(foodRepository)
	foodHandler := _handlerFoods.NewHandler(foodService)

	historiesRepository := _repositoryHistories.NewRepositoryMySQL(db)
	historiesService := _serviceHistories.NewService(historiesRepository, foodRepository, userService, calorieService)
	historiesHandler := _handlerHistories.NewHandler(historiesService)

	// initial of routes
	routesInit := routes.HandlerList{
		JWTMiddleware:    configJWT.Init(),
		UserHandler:      *usersHandler,
		CalorieHandler:   *calorieHandler,
		FoodHandler:      *foodHandler,
		HistoriesHandler: *historiesHandler,
	}
	routesInit.RouteRegister(e)
	_middleware.LogMiddleware(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
