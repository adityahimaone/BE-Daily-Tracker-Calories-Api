package main

import (
	_middleware "daily-tracker-calories/app/middleware"
	_handlerUsers "daily-tracker-calories/app/presenter/users"
	"daily-tracker-calories/app/routes"
	_serviceUsers "daily-tracker-calories/bussiness/users"
	mysqlRepo "daily-tracker-calories/repository/mysql"
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

	db := configDB.IntialDB()
	mysqlRepo.MigrateDB(db)

	e := echo.New()

	//factory of domain
	userRepository := _repositoryUsers.NewRepositoryMySQL(db)
	userService := _serviceUsers.NewService(userRepository)
	usersHandler := _handlerUsers.NewHandler(userService)

	// initial of routes
	routesInit := routes.HandlerList{
		UserHandler: *usersHandler,
	}
	routesInit.RouteRegister(e)
	_middleware.LogMiddleware(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
