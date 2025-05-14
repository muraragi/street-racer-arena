package middleware

import (
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser middleware loads the current user from the session
// and adds it to the context
func CurrentUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Next()
			return
		}

		// Convert userID to uint
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
