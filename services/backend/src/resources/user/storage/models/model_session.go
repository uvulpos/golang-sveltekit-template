package models

import (
	"database/sql"
	"time"

	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

type SessionModel struct {
	ID            string         `db:"id"`
	UserID        string         `db:"user_id"`
	UserAgentHash sql.NullString `db:"useragent_hash"`
	Created       time.Time      `db:"created"`
	CreatedIP     sql.NullString `db:"created_ip_addr"`
	LastRefresh   time.Time      `db:"last_jwt_refresh"`
	LastRefreshIP sql.NullString `db:"last_jwt_refresh_ip_addr"`
}

func (m *SessionModel) ToServiceModel() *serviceModel.SessionModel {

	var hash string
	if m.UserAgentHash.Valid {
		hash = m.UserAgentHash.String
	}

	var createdIP string
	if m.CreatedIP.Valid {
		createdIP = m.CreatedIP.String
	}

	var refreshIP string
	if m.LastRefreshIP.Valid {
		refreshIP = m.LastRefreshIP.String
	}

	return &serviceModel.SessionModel{
		ID:            m.ID,
		UserID:        m.UserID,
		UserAgentHash: hash,
		Created:       m.Created,
		CreatedIP:     createdIP,
		LastRefresh:   m.LastRefresh,
		LastRefreshIP: refreshIP,
	}
}
