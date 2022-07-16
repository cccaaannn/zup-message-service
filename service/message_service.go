package service

import (
	"zup-message-service/database"
	"zup-message-service/model"
)

func GetUnreadMessages(fromId uint64, toId uint64) []model.Message {
	var messages []model.Message
	database.Connection.Where("from_id=? AND to_id=? AND message_status=0", fromId, toId).Find(&messages)
	return messages
}

func GetAllMessages(fromId uint64, toId uint64) []model.Message {
	var messages []model.Message
	database.Connection.Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?)", fromId, toId, toId, fromId).Find(&messages)
	return messages
}

func GetUnReadMessagesAfter(messageId uint64, toId uint64) []model.Message {
	var messages []model.Message
	database.Connection.Where("id>? AND to_id=? AND message_status=0", messageId, toId).Find(&messages)
	return messages
}

func SetMessageAsRead(messageId uint64) bool {
	var message model.Message
	database.Connection.First(&message, messageId)
	message.MessageStatus = 1
	database.Connection.Save(&message)
	return true
}
