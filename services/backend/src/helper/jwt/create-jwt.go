package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	userService "github.com/uvulpos/go-svelte/src/resources/users/service"
)

func NewJWT(userUuid string, roles userService.UserPermissions) (string, error) {
	var key []byte // *ecdsa.PrivateKey  ->> https://golang-jwt.github.io/jwt/usage/create/
	var t *jwt.Token

	var roleIndicators []string
	for _, r := range roles {
		roleIndicators = append(roleIndicators, r.Identifier)
	}

	key = []byte("loafofbread")
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user-uuid": userUuid,
		"roles":     roleIndicators,
	})
	return t.SignedString(key)

}
