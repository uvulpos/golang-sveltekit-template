package storage

import (
	"errors"

	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUserByUsernameOrEmail(identifier string) (*service.UserWithPermission, error) {
	rows := h.db.QueryRow(
		"SELECT * from public.full_user_with_permission WHERE auth_source='basic' AND (username=$1 OR email=$2) LIMIT 1",
		identifier,
		identifier,
	)
	if rows == nil {
		return nil, errors.New("no user found")
	}

	user := UserWithPermission{}
	scanErr := rows.Scan(&user)
	if scanErr != nil {
		return nil, scanErr
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
