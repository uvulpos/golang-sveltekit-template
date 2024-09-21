package models

import "time"

type SessionModel struct {
	ID            string
	UserID        string
	UserAgentHash string
	Created       time.Time
	CreatedIP     string
	LastRefresh   time.Time
	LastRefreshIP string
}
