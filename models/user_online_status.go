package models

import (
	"time"
)

type UserOnlineStatus struct {
	UserId     int       `gorm:"primary_key" json:"user_id"`
	IsOnline   int       `json:"is_online"`
	LastOnline time.Time `json:"last_online"`
}
