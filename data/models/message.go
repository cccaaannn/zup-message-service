package models

import (
	"time"
)

type Message struct {
	ID            uint64    `gorm:"primary_key" json:"id"`
	FromId        uint64    `json:"fromId"`
	ToId          uint64    `json:"toId"`
	MessageText   string    `json:"messageText"`
	MessageStatus int       `json:"messageStatus"`
	MessageType   string    `json:"messageType"`
	CreatedAt     time.Time `json:"createdAt"`
	ReadAt        time.Time `json:"readAt"`
}
