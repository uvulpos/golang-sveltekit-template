package storage

func (h *UserStore) GetUserByUUID(uuid string) (*UserWithPermissions, error) {
	rows, rowsErr := h.db.Query("SELECT * from public.full_user_with_permission")
	if rowsErr != nil {
		return nil, rowsErr
	}

	user := &UserWithPermissions{}
	scanErr := rows.Scan(user)
	if scanErr != nil {
		return nil, scanErr
	}

	return user, nil
}
