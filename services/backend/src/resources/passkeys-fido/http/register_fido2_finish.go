package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FidoRegisterFinishPayload struct {
	ID       string `json:"id"`
	RawID    string `json:"rawId"`
	Type     string `json:"type"`
	Response struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	} `json:"response"`
}

func (h *PasskeyHandler) FinishWebAuthNRegistration(c *fiber.Ctx) error {
	// requestUser := c.Locals("user-uuid")
	// userUuid, ok := requestUser.(string)
	// if !ok {
	// 	return c.SendStatus(http.StatusUnauthorized)
	// }

	// err := h.service.RegisterUserFidoFinish(userUuid)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	return c.SendStatus(http.StatusOK)
}
