package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	httpModels "github.com/uvulpos/golang-sveltekit-template/src/resources/user/http/http-models"
	sessionlocals "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http/consts/session-locals"
)

// @Summary		Get own user permissions
// @Description	Returns all set permission for authorized users
//
// @Tags			user
// @Produce		json
//
// @Success		200	{object}	httpModels.SelfInformationModel
// @Failure		404	{string}	string	"The requested data could not be found."
// @Failure		500	{string}	string	"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/self [get]
func (s *UserHandler) GetSelfInformation(c *fiber.Ctx) error {
	userID := c.Locals(sessionlocals.UserUUID).(string)

	selfInformation, selfInformationErr := s.service.GetUserSelfInformationByID(userID)
	if selfInformationErr != nil {
		status, _, message := selfInformationErr.HttpError()
		return c.Status(status).SendString(message)
	}

	httpModel := httpModels.NewSelfInformationModel(selfInformation)
	return c.Status(http.StatusOK).JSON(httpModel)
}
