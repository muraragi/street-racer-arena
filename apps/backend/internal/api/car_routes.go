package api

import (
	"muraragi/street-racing-arena-backend/internal/handlers"
	"muraragi/street-racing-arena-backend/internal/middleware"
	"muraragi/street-racing-arena-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupCarRoutes(router *gin.Engine, carService services.CarService) {
	carHandler := handlers.NewCarHandler(carService)

	carRoutes := router.Group("/cars")
	carRoutes.Use(middleware.AuthProvider())
	carRoutes.Use(middleware.Auth())
	{
		carRoutes.POST("", carHandler.AddCar)
		carRoutes.GET("", carHandler.GetCars)
		carRoutes.PUT("/:carId", carHandler.UpdateCar)
		carRoutes.DELETE("/:carId", carHandler.DeleteCar)
	}
}
