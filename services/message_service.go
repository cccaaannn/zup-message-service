package services

import (
	"gorm.io/gorm"
	"bytes"
	"encoding/json"
	"zup-message-service/data/dtos"
	"zup-message-service/data/models"
	"zup-message-service/data/enums"
	"zup-message-service/utils/database"
	"zup-message-service/utils/rabbitmq"
)

// func GetUnreadMessages(toId uint64, tokenPayload *dtos.TokenPayload) dtos.DataResult[[]models.Message] {
// 	var messages []models.Message
// 	database.Connection.Where("from_id=? AND to_id=? AND message_status=0", tokenPayload.Id, toId).Find(&messages)
// 	return dtos.DataResult[[]models.Message]{Status: true, Message: "", Data: &messages}
// }

// func GetUnReadMessagesAfter(messageId uint64, toId uint64) dtos.DataResult[[]models.Message] {
// 	var messages []models.Message
// 	database.Connection.Where("id>? AND to_id=? AND message_status=0", messageId, toId).Find(&messages)
// 	return dtos.DataResult[[]models.Message]{Status: true, Message: "", Data: &messages}
// }

func GetConversation(toId uint64, pagination dtos.Pagination, tokenPayload *dtos.TokenPayload) dtos.DataResult[dtos.Pagination] {
	var messages []models.Message
	// database.Connection.Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?) ORDER BY id", tokenPayload.Id, toId, toId, tokenPayload.Id).Find(&messages)

	tx := database.Connection.Model(messages).Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?)", tokenPayload.Id, toId, toId, tokenPayload.Id)
	tx.Scopes(database.Paginate(&pagination, tx)).Find(&messages)
	pagination.Content = messages

	return dtos.DataResult[dtos.Pagination]{Status: true, Message: "", Data: &pagination}
}

func SetMessageAsRead(messageId uint64, tokenPayload *dtos.TokenPayload) dtos.Result {
	var messages []models.Message
	database.Connection.Where("id=? AND from_id=?", messageId, tokenPayload.Id).Find(&messages)

	if len(messages) == 0 {
		return dtos.Result{Status: false, Message: "Message is not belongs to user."}
	}

	message := messages[0]
	message.MessageStatus = 1
	database.Connection.Save(&message)
	return dtos.Result{Status: true, Message: "Message status updated."}
}

func SetMessagesAsRead(userId uint64, tokenPayload *dtos.TokenPayload) dtos.Result {
	database.Connection.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.Message{}).Where("from_id=? AND to_id=?", userId, tokenPayload.Id).Update("MessageStatus", 1)
	return dtos.Result{Status: true, Message: "All messages on this conversation set as read."}
}

func CreateMessage(message *models.Message, accessToken string, tokenPayload *dtos.TokenPayload) dtos.Result {

	message.MessageStatus = 0
	userOnlineStatusResult := GetUserOnlineStatus(message.ToId, accessToken)
	if userOnlineStatusResult.Status && userOnlineStatusResult.Data.OnlineStatus == enums.USER_ONLINE {
		// Status is set to 1 since if user is online it will be directly sent.
		message.MessageStatus = 1
	}

	// From id is used from token for security
	message.FromId = tokenPayload.Id
	message.MessageType = "TEXT"

	// Fist save entity to db to get id
	database.Connection.Create(&message)

	// Publish to queue if user is connected
	if  userOnlineStatusResult.Status && userOnlineStatusResult.Data.OnlineStatus == enums.USER_ONLINE {
		byteBuffer := new(bytes.Buffer)
		json.NewEncoder(byteBuffer).Encode(message)
		rabbitmq.PublishMessage(byteBuffer.Bytes(), rabbitmq.GetQueueNameForUser(message.ToId))

		return dtos.Result{Status: true, Message: "Message created and queued."}
	}

	return dtos.Result{Status: true, Message: "Message created."}
}
