package api

import (
	"muraragi/street-racer-arena-backend/internal/handlers"
	"muraragi/street-racer-arena-backend/internal/middleware"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userService services.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthProvider())
	{
		userRoutes.GET("/", userHandler.UserInfo)
	}
}
