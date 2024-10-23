package service

type AuthentikUserInfo struct {
	ID                string   `json:"sub"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	Name              string   `json:"name"`
	DisplayName       string   `json:"given_name"`
	PreferredUsername string   `json:"preferred_username"`
	Nickname          string   `json:"nickname"`
	Groups            []string `json:"groups"`
}
