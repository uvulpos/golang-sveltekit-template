package service

import "github.com/uvulpos/go-svelte/src/resources/users/storage"

type UserSvc struct {
	storage *storage.UserStore
}

func NewUserSvc(storage *storage.UserStore) *UserSvc {
	return &UserSvc{
		storage,
	}
}
