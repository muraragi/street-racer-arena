package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var (
	googleClientID     string
	googleClientSecret string
	googleCallbackURL  string
	sessionSecret      string
	sessionStore       sessions.Store
)

func InitializeAuth() {
	googleClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	googleClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	googleCallbackURL = os.Getenv("GOOGLE_OAUTH_CALLBACK_URL")
	sessionSecret = os.Getenv("SESSION_SECRET")

	if googleClientID == "" || googleClientSecret == "" || googleCallbackURL == "" || sessionSecret == "" {
		log.Fatal("Auth environment variables not set (GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_CALLBACK_URL, SESSION_SECRET)")
	}

	sessionStore = cookie.NewStore([]byte(sessionSecret))
	sessionStore.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   os.Getenv("GIN_MODE") == "release",
		SameSite: http.SameSiteLaxMode,
	})
	gothic.Store = sessionStore

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, googleCallbackURL, "email", "profile"),
	)
}

func GetSessionStore() sessions.Store {
	return sessionStore
}
