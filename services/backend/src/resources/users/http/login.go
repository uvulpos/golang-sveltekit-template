package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/helper/jwt"
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

	user, authErr := h.service.GetAuthenticatedUserByCredentials(nil, payload.Username, payload.Password)
	if authErr != nil {
		return c.Status(http.StatusUnauthorized).SendString(authErr.Error())
	}

	token, tokenErr := jwt.NewJWT(user)
	if tokenErr != nil {
		return errors.New("internal server error")
	}

	c.SendString(token)

	return nil
}
