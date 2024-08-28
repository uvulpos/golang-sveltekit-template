package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	httpModels "github.com/uvulpos/golang-sveltekit-template/src/resources/user/http/http-models"
)

// @Summary		Get own user permissions
// @Description	Returns all set permission for authorized users
//
// @Tags			user
// @Produce		plain
//
// @Success		200	{object}	httpModels.SelfPermissionsModel
// @Failure		404	{string}	string	"The requested data could not be found."
// @Failure		500	{string}	string	"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/self/permissions [get]
func (s *UserHandler) GetSelfPermissions(c *fiber.Ctx) error {
	userID := c.Locals("user-uuid").(string)

	permissions, permissionsErr := s.service.GetUserPermissionsByID(userID)
	if permissionsErr != nil {
		status, _, message := permissionsErr.HttpError()
		return c.Status(status).SendString(message)
	}

	permissionModel := httpModels.NewSelfPermissionsModel(permissions)
	return c.Status(http.StatusOK).JSON(permissionModel)
}
