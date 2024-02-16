package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *PasskeyHandler) BeginWebAuthNRegistration(c *fiber.Ctx) error {
	// requestUser := c.Locals("user-uuid")
	// userUuid, ok := requestUser.(string)
	// if !ok {
	// 	return c.SendStatus(http.StatusUnauthorized)
	// }

	// credential, err := h.service.RegisterUserFidoBegin(userUuid)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// return c.JSON(credential)

	return c.SendStatus(http.StatusOK)
}
