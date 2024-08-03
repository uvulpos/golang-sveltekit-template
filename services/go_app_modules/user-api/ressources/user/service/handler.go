package service

type UserService struct {
	storage UserStore
}

func NewUserService(storage UserStore) *UserService {
	return &UserService{
		storage,
	}
}

type UserStore interface {
}
