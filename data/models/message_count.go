package models

type MessageCount struct {
	FromId uint64 `json:"fromId"`
	Count  uint64 `json:"count"`
}
