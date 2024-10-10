package service

import (
	"fmt"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func HandlePanic() {
	r := recover()

	if r != nil {
		fmt.Println("RECOVER", r)
	}
}

func (s *UserService) GetUserPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface) {
	defer HandlePanic()
	return s.storage.GetUserPermissionsByID(tx, userID)
}
