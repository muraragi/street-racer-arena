package handlers

import (
	"net/http"
	"strconv"

	"muraragi/street-racing-arena-backend/internal/models"
	"muraragi/street-racing-arena-backend/internal/services"
	"muraragi/street-racing-arena-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type CarHandler interface {
	AddCar(c *gin.Context)
	GetCars(c *gin.Context)
	UpdateCar(c *gin.Context)
	DeleteCar(c *gin.Context)
}

type carHandler struct {
	carService services.CarService
}

func NewCarHandler(carService services.CarService) CarHandler {
	return &carHandler{carService: carService}
}

func (h *carHandler) AddCar(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)
	if currentUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	input := &models.CarDTO{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := &models.Car{
		UserID:         currentUser.ID,
		BaseCarModelID: input.BaseCarModelID,
		Nickname:       input.Nickname,
	}

	createdCar, err := h.carService.AddCarToUser(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCar)
}

func (h *carHandler) GetCars(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)

	cars, err := h.carService.GetUsersCars(currentUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (h *carHandler) UpdateCar(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)

	carIDParam := c.Param("carId")
	carID, err := strconv.ParseUint(carIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car_id"})
		return
	}

	input := &models.CarDTO{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := &models.Car{
		UserID:         currentUser.ID,
		BaseCarModelID: input.BaseCarModelID,
		Nickname:       input.Nickname,
	}
	car.ID = uint(carID)

	updatedCar, err := h.carService.UpdateUserCar(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCar)
}

func (h *carHandler) DeleteCar(c *gin.Context) {
	carIDParam := c.Param("carId")
	carID, err := strconv.ParseUint(carIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car_id"})
		return
	}

	if err := h.carService.DeleteUserCar(uint(carID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
