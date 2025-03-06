package api

import (
	"bookstore-framework/internal/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	router.POST("/register", userHandler.RegisterHandler)
}
