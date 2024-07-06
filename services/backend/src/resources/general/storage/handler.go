package storage

import (
	"github.com/go-sqlx/sqlx"
)

type GeneralStore struct {
	dbstore *sqlx.DB
}

func NewGeneralStore(db *sqlx.DB) *GeneralStore {
	return &GeneralStore{
		db,
	}
}
