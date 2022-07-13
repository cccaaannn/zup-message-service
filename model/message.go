package model

import (
	"time"
)

type Message struct {
	ID            uint64    `gorm:"primary_key" json:"id"`
	FromId        int       `json:"from_id"`
	ToId          int       `json:"to_id"`
	MessageText   string    `json:"message_text"`
	MessageStatus int       `json:"message_status"`
	CreatedAt     time.Time `json:"created_at"`
}
