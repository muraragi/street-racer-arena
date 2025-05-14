package api

import (
	"muraragi/street-racer-arena-backend/internal/auth"
	"muraragi/street-racer-arena-backend/internal/middleware"
	"muraragi/street-racer-arena-backend/internal/services"
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
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-CSRF-Token", "Your-Custom-Header-If-Any"}, // Be very broad
		ExposeHeaders:    []string{"Content-Length", "street_racer_session", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

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
