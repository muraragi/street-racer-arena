package middleware

import (
	"github.com/gin-gonic/gin"
)

func GoogleAuthProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()

		if q.Get("provider") == "" {
			q.Set("provider", "google")
			c.Request.URL.RawQuery = q.Encode()
		}

		c.Next()
	}
}
