package storage

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, customerrors.NewDatabaseTransactionBeginError(err, "Failed to start transaction")
	}
	return tx, nil
}
