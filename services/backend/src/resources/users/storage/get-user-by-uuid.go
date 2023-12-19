package storage

import (
	"errors"

	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUserByUUID(uuid string) (*service.UserWithPermission, error) {
	rows := h.dbstore.DB.QueryRowx("SELECT * from public.full_user_with_permission WHERE id=$1 LIMIT 1", uuid)
	if rows == nil {
		return nil, errors.New("no user found")
	}

	user := &UserWithPermission{}
	scanErr := rows.StructScan(user)
	if scanErr != nil {
		return nil, scanErr
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
