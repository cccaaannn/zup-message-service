package services

import (
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

func GetAllMessages(toId uint64, tokenPayload *dtos.TokenPayload) dtos.DataResult[[]models.Message] {
	var messages []models.Message
	database.Connection.Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?) ORDER BY id", tokenPayload.Id, toId, toId, tokenPayload.Id).Find(&messages)
	return dtos.DataResult[[]models.Message]{Status: true, Message: "", Data: &messages}
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

func CreateMessage(message *models.Message, accessToken string, tokenPayload *dtos.TokenPayload) dtos.Result {

	// From id is used from token for security
	message.FromId = tokenPayload.Id
	message.MessageStatus = 0

	// Fist save entity to db to get id
	database.Connection.Create(&message)

	userOnlineStatusResult := GetUserOnlineStatus(message.ToId, accessToken)

	// Publish to queue if user is connected
	if userOnlineStatusResult.Status && userOnlineStatusResult.Data.OnlineStatus == enums.USER_ONLINE {
		byteBuffer := new(bytes.Buffer)
		json.NewEncoder(byteBuffer).Encode(message)
		rabbitmq.PublishMessage(byteBuffer.Bytes(), rabbitmq.GetQueueNameForUser(message.ToId))

		return dtos.Result{Status: true, Message: "Message created and queued."}
	}

	return dtos.Result{Status: true, Message: "Message created."}
}
