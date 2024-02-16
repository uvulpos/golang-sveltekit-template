package storage

import (
	"errors"
	"time"

	"github.com/go-sqlx/sqlx"
)

func (h *PasskeyStore) InsertFidoSession(tx *sqlx.Tx, uuid string, sessionJson []byte, expires time.Time) error {
	var rowErr error
	const sql = ``

	if tx == nil {
		_, rowErr = h.dbstore.DB.Exec(sql)
	} else {
		_, rowErr = tx.Exec(sql)
	}

	if rowErr != nil {
		return errors.New("could not insert session into database")
	}

	return nil
}
