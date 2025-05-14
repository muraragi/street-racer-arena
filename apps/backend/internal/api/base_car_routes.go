package api

import (
	"muraragi/street-racing-arena-backend/internal/handlers"
	"muraragi/street-racing-arena-backend/internal/middleware"
	"muraragi/street-racing-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupBaseCarRoutes(router *gin.Engine, baseCarService services.BaseCarService) {
	baseCarHandler := handlers.NewBaseCarHandler(baseCarService)

	baseCarRoutes := router.Group("/base-car")
	baseCarRoutes.Use(middleware.Auth())
	{
		baseCarRoutes.GET("/", baseCarHandler.GetAllBaseCars)
	}
}
