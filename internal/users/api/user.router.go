package api

import (
	"bookstore-framework/internal/users"
	"bookstore-framework/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	router.POST("/register", userHandler.RegisterHandler)
	router.POST("/login", userHandler.LoginHandler)

	protected := router.Group("/")
	protected.Use(middleware.JWTAuth())
	protected.GET("/profile", userHandler.GetProfile)
}
