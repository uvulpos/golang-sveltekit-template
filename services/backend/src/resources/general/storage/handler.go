package storage

import (
	"github.com/uvulpos/go-svelte/src/helper/database"
)

type GeneralStore struct {
	dbstore database.Sql
}

func NewGeneralStore(db database.Sql) *GeneralStore {
	return &GeneralStore{
		db,
	}
}
