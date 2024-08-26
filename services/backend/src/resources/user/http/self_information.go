package http

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get own user permissions
// @Description	Returns all set permission for authorized users
//
// @Tags			user
// @Produce		plain
//
// @Success		200	{object}	httpModels.SelfPermissions	"{"Permissions":["greetings:read","greetings:write"]}"
// @Failure		404	{string}	string						"The requested data could not be found."
// @Failure		500	{string}	string						"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/self/permissions [get]
func (s *UserHandler) GetSelfInformation(c *fiber.Ctx) error {
	return nil
	// userID := c.Locals("user-uuid").(string)

	// permissions, permissionsErr := s.service.GetUserSelfInformationByID(userID)
	// if permissionsErr != nil {
	// 	status, _, message := permissionsErr.HttpError()
	// 	return c.Status(status).SendString(message)
	// }

	// permissionModel := httpModels.NewSelfPermissions(permissions)

	// jsonPermissions, jsonPermissionsErr := permissionModel.ToJson()
	// if jsonPermissionsErr != nil {
	// 	status, _, message := jsonPermissionsErr.HttpError()
	// 	return c.Status(status).SendString(message)
	// }

	// return c.Status(http.StatusOK).SendString(jsonPermissions)
}
