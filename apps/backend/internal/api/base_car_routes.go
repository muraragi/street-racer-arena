package api

import (
	"muraragi/street-racer-arena-backend/internal/handlers"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupBaseCarRoutes(router *gin.Engine, baseCarService services.BaseCarService) {
	baseCarHandler := handlers.NewBaseCarHandler(baseCarService)

	baseCarRoutes := router.Group("/base-car")
	{
		baseCarRoutes.GET("/", baseCarHandler.GetAllBaseCars)
	}
}
