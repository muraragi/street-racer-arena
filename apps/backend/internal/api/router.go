package api

import (
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-contrib/cors"
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

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	SetupBaseCarRoutes(router, dependencies.BaseCarService)
	SetupAuthRoutes(router, dependencies.UserService)

	return router
}
