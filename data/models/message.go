package models

import (
	"time"
)

type Message struct {
	ID            uint64    `gorm:"primary_key" json:"id"`
	FromId        uint64    `json:"from_id"`
	ToId          uint64    `json:"to_id"`
	MessageText   string    `json:"message_text"`
	MessageStatus int       `json:"message_status"`
	CreatedAt     time.Time `json:"created_at"`
}
