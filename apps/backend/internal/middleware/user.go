package middleware

import (
	"muraragi/street-racing-arena-backend/internal/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Next()
			return
		}

		id, ok := userID.(uint)
		if !ok {
			c.Next()
			return
		}

		user, err := userService.GetUserByID(id)
		if err != nil {
			c.Next()
			return
		}

		c.Set("current_user", user)
		c.Next()
	}
}
