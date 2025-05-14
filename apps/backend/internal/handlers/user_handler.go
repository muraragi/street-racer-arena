package handlers

import (
	"muraragi/street-racer-arena-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, _ := h.userService.GetUserFromSession(gothUser)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
