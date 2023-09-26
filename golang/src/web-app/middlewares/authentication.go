package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	// jwtToken := c.Cookies("jwt", "")

	// if jwtToken == "" {
	// 	return errors.New("no authentication token found")
	// }

	// token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte("my-jwt-password"), nil
	// })

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	c.Locals("user-uuid", claims["userid"])
	// } else {
	// 	c.Status(403).SendString("jwt is not valid")
	// 	return errors.New("jwt is not valid")
	// }

	return nil
}
