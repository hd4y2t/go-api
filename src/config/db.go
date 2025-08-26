package config

import (
	"gorm.io/gorm",
	"gorm.io/driver/mysql"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := 3306
	user := "root"
	password := ""
	dbname := "go-api"

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}