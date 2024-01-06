package storage

import (
	"errors"

	"github.com/uvulpos/go-svelte/src/resources/users/service"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserStore) GetUserByCredentials(username, password string) (*service.UserWithPermission, error) {
	rows := h.dbstore.DB.QueryRowx(
		`SELECT * from public.full_user_with_permission 
		WHERE auth_source='basic' AND (username=$1 OR email=$2) LIMIT 1;`,
		username,
		username,
	)
	if rows == nil {
		return nil, errors.New("no user found")
	}

	user := UserWithPermission{}
	scanErr := rows.StructScan(&user)
	if scanErr != nil {
		return nil, scanErr
	}

	userPasswordVerifyErr := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if userPasswordVerifyErr != nil || user.Password.String == "" || password == "" {
		return nil, errors.New("no user found")
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
