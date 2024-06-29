package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/uvulpos/go-svelte/src/configuration"
	userService "github.com/uvulpos/go-svelte/src/resources/users/service"
)

func NewJWT(user *userService.UserWithPermission) (string, error) {
	var key []byte // *ecdsa.PrivateKey  ->> https://golang-jwt.github.io/jwt/usage/create/
	var t *jwt.Token

	var roles userService.UserPermissions = user.Permissions
	var roleIndicators []string
	for _, r := range roles {
		roleIndicators = append(roleIndicators, r.Identifier)
	}

	key = []byte(configuration.Configuration.Session.JwtSecret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user-uuid":   user.Id.String(),
		"username":    user.Username,
		"roleName":    user.Role.Name,
		"permissions": roleIndicators,
	})
	return t.SignedString(key)

}
