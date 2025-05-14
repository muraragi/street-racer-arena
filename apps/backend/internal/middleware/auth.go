package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := gothic.CompleteUserAuth(c.Writer, c.Request)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
