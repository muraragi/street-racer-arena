package api

import (
	"muraragi/street-racer-arena-backend/internal/handlers"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userService services.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", userHandler.UserInfo)
	}
}
