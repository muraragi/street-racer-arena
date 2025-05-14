package api

import (
	"muraragi/street-racing-arena-backend/internal/handlers"
	"muraragi/street-racing-arena-backend/internal/middleware"
	"muraragi/street-racing-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userService services.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.GoogleAuthProvider())
	userRoutes.Use(middleware.AuthProvider())
	userRoutes.Use(middleware.Auth())
	{
		userRoutes.GET("/", userHandler.UserInfo)
	}
}
