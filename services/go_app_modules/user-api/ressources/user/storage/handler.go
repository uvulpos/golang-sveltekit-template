package storage

import "github.com/go-sqlx/sqlx"

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{
		db,
	}
}
