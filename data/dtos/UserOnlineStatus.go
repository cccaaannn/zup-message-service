package dtos

import "time"

type UserOnlineStatus struct {
	Id           uint64    `json:"id"`
	OnlineStatus string    `json:"onlineStatus"`
	LastOnline   time.Time `json:"lastOnline"`
}
