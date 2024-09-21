package models

import (
	"time"

	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

type SessionModel struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	UserAgentHash string    `db:"useragent_hash"`
	Created       time.Time `db:"created"`
	CreatedIP     string    `db:"created_ip_addr"`
	LastRefresh   time.Time `db:"last_jwt_refresh"`
	LastRefreshIP string    `db:"last_jwt_refresh_ip_addr"`
}

func (m *SessionModel) ToServiceModel() *serviceModel.SessionModel {
	return &serviceModel.SessionModel{
		ID:            m.ID,
		UserID:        m.UserID,
		UserAgentHash: m.UserAgentHash,
		Created:       m.Created,
		CreatedIP:     m.CreatedIP,
		LastRefresh:   m.LastRefresh,
		LastRefreshIP: m.LastRefreshIP,
	}
}
