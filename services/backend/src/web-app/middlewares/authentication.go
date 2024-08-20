package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
)

func Authentication(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt", "")
	authorizationHeader := c.GetReqHeaders()["Authorization"]

	if jwtToken == "" && len(authorizationHeader) > 0 {
		jwtToken = authorizationHeader[0]
	}

	jwtData, jwtDataErr := jwtService.VerifyJWT(jwtToken, "somethingNice")
	if jwtDataErr != nil {
		fmt.Println("JWT ERROR", jwtDataErr)
		return c.Status(fiber.StatusInternalServerError).SendString("could not verify authentication")
	}

	// after jwt got evaluated, add information from jwt to fiber request context
	c.Locals("jwt-token", jwtData.UserUuid)

	return c.Next()
}
