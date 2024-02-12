package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *PasskeyHandler) FinishWebAuthNRegistration(c *fiber.Ctx) error {
	fmt.Println("Verifiy FIDO 2 Challenge")

	h.service.RegisterUserFidoFinish()
	return c.SendStatus(http.StatusOK)
}
