package config

import (
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := 3306
	user := "root"
	password := ""
	dbname := "go-api"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}