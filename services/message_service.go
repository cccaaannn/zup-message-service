package services

import (
	"strconv"

	"bytes"
	"encoding/json"
	"zup-message-service/database"
	"zup-message-service/models"
	"zup-message-service/rabbitmq"
)

func GetUnreadMessages(fromId uint64, toId uint64) []models.Message {
	var messages []models.Message
	database.Connection.Where("from_id=? AND to_id=? AND message_status=0", fromId, toId).Find(&messages)
	return messages
}

func GetAllMessages(fromId uint64, toId uint64) []models.Message {
	var messages []models.Message
	database.Connection.Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?) ORDER BY id", fromId, toId, toId, fromId).Find(&messages)
	return messages
}

func GetUnReadMessagesAfter(messageId uint64, toId uint64) []models.Message {
	var messages []models.Message
	database.Connection.Where("id>? AND to_id=? AND message_status=0", messageId, toId).Find(&messages)
	return messages
}

func SetMessageAsRead(messageId uint64) bool {
	var message models.Message
	database.Connection.First(&message, messageId)
	message.MessageStatus = 1
	database.Connection.Save(&message)
	return true
}

func CreateMessage(message *models.Message, accessToken string) bool {

	// Fist save entity to db to get id
	database.Connection.Create(&message)

	userOnlineStatusResult := GetUserOnlineStatus(message.ToId, accessToken)

	// Publish to queue if user is connected
	if userOnlineStatusResult.Status && userOnlineStatusResult.Data.OnlineStatus == "ONLINE" {
		byteBuffer := new(bytes.Buffer)
		json.NewEncoder(byteBuffer).Encode(message)
		rabbitmq.PublishMessage(byteBuffer.Bytes(), strconv.FormatUint(message.ToId, 10))
	}

	return true
}
