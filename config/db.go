package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	password := ""
	dbname := "go_api"

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}