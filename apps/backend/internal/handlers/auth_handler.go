package handlers

import (
	"fmt"
	"net/http"

	"muraragi/street-racing-arena-backend/internal/services"

	"github.com/gin-contrib/sessions"
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

	if _, authErr := gothic.CompleteUserAuth(c.Writer, c.Request); authErr == nil {
		if redirectURL == "" {
			redirectURL = "https://street-racing-arena.muraragi.com"
		}

		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
		return
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

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("provider", gothUser.Provider)
	session.Set("provider_id", gothUser.UserID)
	session.Set("username", user.Username)
	session.Set("avatar_url", user.AvatarURL)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, "https://street-racing-arena.muraragi.com")
}

func (h *authHandler) Logout(c *gin.Context) {
	redirectURL := c.Query("redirect_url")

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	gothic.Logout(c.Writer, c.Request)

	if redirectURL == "" {
		redirectURL = "https://street-racing-arena.muraragi.com"
	}

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
