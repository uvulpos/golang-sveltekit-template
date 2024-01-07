package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	userService "github.com/uvulpos/go-svelte/src/resources/users/service"
)

const jwtTokenSecret = `loafofbread`

func NewJWT(user *userService.UserWithPermission) (string, error) {
	var key []byte // *ecdsa.PrivateKey  ->> https://golang-jwt.github.io/jwt/usage/create/
	var t *jwt.Token

	var roles userService.UserPermissions = user.Permissions
	var roleIndicators []string
	for _, r := range roles {
		roleIndicators = append(roleIndicators, r.Identifier)
	}

	key = []byte("loafofbread")
	t = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user-uuid":   user.Id.String(),
		"username":    user.Username,
		"roleName":    user.RoleName,
		"permissions": roleIndicators,
	})
	return t.SignedString(key)

}
