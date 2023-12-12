package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUserByCredentials(username, password string) (*service.UserWithPermission, error) {
	// Mock due to problems with the sql driver, not happy about it yet
	user := UserWithPermission{
		Id:                  uuid.New(),
		Username:            "uvulpos",
		Email:               "uvulpos@home.de",
		Password:            "123",
		LdapUUID:            "",
		AuthSource:          "basic",
		AdminReviewRequired: false,
		RoleID:              "09e7a7ab-0895-4fac-a57f-591b3a1340b5",
		RoleName:            "Admin",
		Permissions: UserPermissions{
			{
				ID:          uuid.New(),
				Name:        "Admin Greet",
				Description: "A unique greeting that just an admin gets",
				Identifier:  "GREET_ADMIN",
			},
		},
	}

	if password != user.Password {
		return nil, errors.New("no user found")
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil

	// ---------

	// rows := h.db.QueryRow(
	// 	"SELECT * from public.full_user_with_permission WHERE auth_source='basic' AND (username=$1 OR email=$2) LIMIT 1",
	// 	username,
	// 	username,
	// )
	// if rows == nil {
	// 	fmt.Println("1")
	// 	return nil, errors.New("no user found")
	// }

	// user := UserWithPermission{}
	// scanErr := rows.Scan(&user)
	// if scanErr != nil {
	// 	fmt.Println(scanErr)
	// 	return nil, scanErr
	// }

	// // todo: check password hash
	// if user.Password == "" {
	// 	fmt.Println("3")
	// 	return nil, errors.New("no user found")
	// }

	// svcUser := user.ToUserWithPermissionsSvcModel()
	// return svcUser, nil
}
