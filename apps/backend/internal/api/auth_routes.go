package api

import (
	"muraragi/street-racer-arena-backend/internal/handlers"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, userService services.UserService) {
	authHandler := handlers.NewAuthHandler(userService)

	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/login", authHandler.BeginAuth)
		authRoutes.GET("/google/callback", authHandler.AuthCallback)
	}
}
