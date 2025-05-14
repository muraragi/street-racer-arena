package handlers

import (
	"muraragi/street-racing-arena-backend/internal/services"
	"muraragi/street-racing-arena-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserInfo(c *gin.Context)
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
