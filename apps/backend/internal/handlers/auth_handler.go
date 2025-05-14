package handlers

import (
	"fmt"
	"muraragi/street-racer-arena-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler interface {
	BeginAuth(c *gin.Context)
	AuthCallback(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) AuthHandler {
	return &authHandler{userService: userService}
}

func (h *authHandler) BeginAuth(c *gin.Context) {
	redirectURL := c.Query("redirect_url")

	if redirectURL == "" {
		redirectURL = "/"
	}

	gothic.StoreInSession("redirect_url", redirectURL, c.Request, c.Writer)

	if gothUser, authErr := gothic.CompleteUserAuth(c.Writer, c.Request); authErr == nil {
		user, err := h.userService.GetUserFromSession(gothUser)

		if err != nil {
			fmt.Fprintln(c.Writer, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logged in", "user": user})
	}

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h *authHandler) AuthCallback(c *gin.Context) {
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	user, err := h.userService.CreateUser(gothUser)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "user": user})

	redirectURL, _ := gothic.GetFromSession("redirect_url", c.Request)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (h *authHandler) Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
