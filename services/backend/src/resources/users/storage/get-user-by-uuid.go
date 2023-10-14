package storage

import "github.com/uvulpos/go-svelte/src/resources/users/service"

func (h *UserStore) GetUserByUUID(uuid string) (*service.UserWithPermission, error) {
	rows, rowsErr := h.dbstore.DB.Query("SELECT * from public.full_user_with_permission WHERE id=$1 LIMIT 1", uuid)
	if rowsErr != nil {
		return nil, rowsErr
	}

	user := &UserWithPermission{}
	scanErr := rows.Scan(user)
	if scanErr != nil {
		return nil, scanErr
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
