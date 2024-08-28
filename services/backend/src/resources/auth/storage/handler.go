package storage

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

type AuthStore struct {
	db *sqlx.DB
}

func NewAuthStore(db *sqlx.DB) *AuthStore {
	return &AuthStore{
		db,
	}
}

func (s *AuthStore) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, customerrors.NewDatabaseTransactionBeginError(err, "Failed to start transaction")
	}
	return tx, nil
}
