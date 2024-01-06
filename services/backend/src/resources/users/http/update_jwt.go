package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/helper/errors"
	"github.com/uvulpos/go-svelte/src/helper/jwt"
)

type RefreshJWTPayload struct {
	Username string `json:"username" xml:"username" form:"username"`
}

func (h *UserHandler) HandleJWTRefresh(c *fiber.Ctx) error {
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	payload := RefreshJWTPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	user, authErr := h.service.GetUserByUUID(nil, userUuid)
	if authErr != nil {
		return c.Status(http.StatusUnauthorized).SendString(authErr.Error())
	}

	token, tokenErr := jwt.NewJWT(user)
	if tokenErr != nil {
		return errors.NewInternalServerErrorApp("could not create jwt").ToHttpError()
	}

	c.SendString(token)
	return nil
}
