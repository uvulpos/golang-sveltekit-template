package service

import (
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (s *UserService) GetUserIDByLogin(providerType providerConst.AuthProvider, providerID string) (string, customerrors.ErrorInterface) {
	return s.storage.GetUserIDByLogin(string(providerType), providerID)
}
