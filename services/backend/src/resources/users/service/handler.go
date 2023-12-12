package service

type UserSvc struct {
	storage UserStorage
}

func NewUserSvc(storage UserStorage) *UserSvc {
	return &UserSvc{
		storage,
	}
}

type UserStorage interface {
	GetUserByCredentials(username, password string) (*UserWithPermission, error)
	GetUserByUUID(uuid string) (*UserWithPermission, error)
}
