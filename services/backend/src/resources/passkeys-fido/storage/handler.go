package storage

import (
	"github.com/uvulpos/go-svelte/src/helper/database"
)

type PasskeyStore struct {
	dbstore database.Sql
}

func NewUserStore(db database.Sql) *PasskeyStore {
	return &PasskeyStore{
		db,
	}
}
