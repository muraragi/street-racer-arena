package api

import (
	"muraragi/street-racer-arena-backend/internal/auth"
	"muraragi/street-racer-arena-backend/internal/middleware"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RouterDependencies struct {
	BaseCarService services.BaseCarService
	UserService    services.UserService
}

func InitializeRouter(dependencies RouterDependencies) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	store := auth.GetSessionStore()

	router.Use(sessions.Sessions("street_racer_session", store))
	router.Use(middleware.CurrentUser(dependencies.UserService))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	SetupBaseCarRoutes(router, dependencies.BaseCarService)
	SetupUserRoutes(router, dependencies.UserService)
	SetupAuthRoutes(router, dependencies.UserService)

	return router
}
