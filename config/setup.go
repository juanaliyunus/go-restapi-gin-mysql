package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase()  {
	database, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/go_restapi_gin_db"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB = database
}