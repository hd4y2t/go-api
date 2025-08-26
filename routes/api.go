package routes

import (
	"github.com/gin-gonic/gin"
	user "go-api/moduls/users"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService, ctx)

	api := router.Group("/api")
	{
		api.GET("/users", userController.GetAllUser)
		api.GET("/users/:id", userController.GetById)
		api.POST("/users", userController.Create)
		api.PUT("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)
	}
}
