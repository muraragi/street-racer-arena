package utils

import (
	"muraragi/street-racer-arena-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) *models.User {
	user, exists := c.Get("current_user")
	if !exists {
		return nil
	}

	currentUser, ok := user.(*models.User)
	if !ok {
		return nil
	}

	return currentUser
}

func IsAuthenticated(c *gin.Context) bool {
	return GetCurrentUser(c) != nil
}
