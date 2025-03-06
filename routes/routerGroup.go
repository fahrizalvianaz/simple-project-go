package routes

import (
	"bookstore-framework/internal/users/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	group := router.Group("/api/v1")

	api.RegisterUserRoutes(group.Group("/users"), db)

	return router
}
