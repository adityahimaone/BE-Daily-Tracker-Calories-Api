package mysql

import (
	"daily-tracker-calories/repository/mysql/calories"
	"daily-tracker-calories/repository/mysql/foods"
	"daily-tracker-calories/repository/mysql/histories"
	"daily-tracker-calories/repository/mysql/users"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func (config *ConfigDB) IntialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&calories.Calories{})
	db.AutoMigrate(&foods.Foods{})
	db.AutoMigrate(&histories.Histories{})
}
