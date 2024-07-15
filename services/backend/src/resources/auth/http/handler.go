package http

import authService "github.com/uvulpos/go-svelte/authentication-api/ressources/auth/service"

type AuthHandler struct {
	auth authService.AuthService
}

func NewAuthHandler(OauthKey, OauthSecret, OauthCallbackURL string) *AuthHandler {
	return &AuthHandler{
		auth: *authService.NewAuthService(OauthKey, OauthSecret, OauthCallbackURL, "email", "username"),
	}
}
