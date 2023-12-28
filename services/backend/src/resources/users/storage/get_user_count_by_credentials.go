package storage

func (h *UserStore) GetUserCountByCredentials(username string) (int, error) {
	rows := h.dbstore.DB.QueryRowx(
		`SELECT COUNT(id) from public.full_user_with_permission 
		WHERE auth_source='basic' AND (username=$1 OR email=$2) LIMIT 1;`,
		username,
		username,
	)

	var userCount int
	scanErr := rows.Scan(&userCount)
	if scanErr != nil {
		return 0, scanErr
	}
	return userCount, nil
}
