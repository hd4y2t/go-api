package main

import (
	"go-api/config"

	"github.com/gin-gonic/gin"

	route "go-api/routes"
)

func main() {
	r := gin.Default()
	db := config.DB()

	route.Api(r, db)

	r.Run()
}
