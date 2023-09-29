package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/helper/cookies"
)

type LoginPayload struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

func (h *UserHandler) HandleLogin(c *fiber.Ctx) error {
	payload := LoginPayload{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	user, userErr := h.service.GetProfileByUsernameOrEmail(payload.Username)
	if userErr != nil {
		c.Status(http.StatusInternalServerError)
		return userErr
	}

	fmt.Println(user)

	jwtToken := fmt.Sprintf("User %s, Pass %s", payload.Username, payload.Password)
	authCookie := cookies.CreateAuthenticationFiberCookie(jwtToken)
	c.Cookie(authCookie)

	return nil
}
