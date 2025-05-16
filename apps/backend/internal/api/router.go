package api

import (
	"muraragi/street-racing-arena-backend/internal/auth"
	"muraragi/street-racing-arena-backend/internal/middleware"
	"muraragi/street-racing-arena-backend/internal/services"
	"time"

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

	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://street-racing-arena.muraragi.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	store := auth.GetSessionStore()

	router.Use(sessions.Sessions("racing_session", store))
	router.Use(middleware.CurrentUser(dependencies.UserService))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	SetupUserRoutes(router, dependencies.UserService)
	SetupAuthRoutes(router, dependencies.UserService)
	SetupBaseCarRoutes(router, dependencies.BaseCarService)

	return router
}
