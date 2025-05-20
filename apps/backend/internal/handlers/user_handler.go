package handlers

import (
	"muraragi/street-racing-arena-backend/internal/services"
	"muraragi/street-racing-arena-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserInfo(c *gin.Context)
	SetSelectedCar(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) UserInfo(c *gin.Context) {
	user := utils.GetCurrentUser(c)

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) SetSelectedCar(c *gin.Context) {
	user := utils.GetCurrentUser(c)

	carIDParam := c.Param("car_id")
	carIDUint, err := strconv.ParseUint(carIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car_id"})
		return
	}
	carID := uint(carIDUint)

	h.userService.SetSelectedCar(user.ID, carID)
}
