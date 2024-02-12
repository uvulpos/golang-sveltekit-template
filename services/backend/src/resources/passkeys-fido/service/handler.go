package service

import (
	"time"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/helper/webauthn"

	userSvc "github.com/uvulpos/go-svelte/src/resources/users/service"
)

type PasskeySvc struct {
	storage     PasskeyStorage
	userService *userSvc.UserSvc
	webAuthN    *webauthn.WebAuthN
}

func NewPasskeySvc(storage PasskeyStorage, user *userSvc.UserSvc, webAuthN *webauthn.WebAuthN) *PasskeySvc {
	return &PasskeySvc{
		storage,
		user,
		webAuthN,
	}
}

type PasskeyStorage interface {
	CreateTransaction() (*sqlx.Tx, error)
	InsertFidoSession(tx *sqlx.Tx, uuid string, sessionJson []byte, expires time.Time) error
}
