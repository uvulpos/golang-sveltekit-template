package service

import (
	"errors"

	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (s *UserService) GetUserIDByLogin(providerType providerConst.AuthProvider, providerID string) (string, customerrors.ErrorInterface) {
	return "", customerrors.NewInternalServerError(errors.New("function not implemented"), "", "function not implemented")
}
