package services

import (
	"log"
	"gorm.io/gorm"
	"bytes"
	"time"
	"encoding/json"
	"zup-message-service/data/dtos"
	"zup-message-service/data/models"
	"zup-message-service/data/enums"
	"zup-message-service/utils/database"
	"zup-message-service/utils/rabbitmq"
)

func GetConversation(toId uint64, pagination dtos.Pagination, tokenPayload *dtos.TokenPayload) dtos.DataResult[dtos.Pagination] {
	var messages []models.Message
	tx := database.Connection.Model(messages).Where("(from_id=? AND to_id=?) OR (from_id=? AND to_id=?)", tokenPayload.Id, toId, toId, tokenPayload.Id)
	tx.Scopes(database.Paginate(&pagination, tx)).Find(&messages)
	pagination.Content = messages

	return dtos.DataResult[dtos.Pagination]{Status: true, Message: "", Data: &pagination}
}

func SetMessageAsRead(messageId uint64, tokenPayload *dtos.TokenPayload) dtos.Result {
	var messages []models.Message
	database.Connection.Where("id=? AND to_id=?", messageId, tokenPayload.Id).Find(&messages)

	if len(messages) == 0 {
		return dtos.Result{Status: false, Message: "Message is not belongs to user."}
	}

	message := messages[0]
	message.MessageStatus = 1
	message.ReadAt = time.Now()
	database.Connection.Save(&message)
	return dtos.Result{Status: true, Message: "Message status updated."}
}

func SetMessagesAsRead(userId uint64, tokenPayload *dtos.TokenPayload) dtos.Result {
	database.Connection.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.Message{}).Where("from_id=? AND to_id=?", userId, tokenPayload.Id).Updates(map[string]interface{}{"MessageStatus": 1, "ReadAt": time.Now()})
	return dtos.Result{Status: true, Message: "All messages on this conversation set as read."}
}

func GetUnreadMessageCount(tokenPayload *dtos.TokenPayload) dtos.ListDataResult[models.MessageCount] {
	var messageCounts []models.MessageCount
	database.Connection.Model(&models.Message{}).Select("from_id, COUNT(*)").Group("from_id").Where("to_id=? AND message_status=0", tokenPayload.Id).Find(&messageCounts)
	return dtos.ListDataResult[models.MessageCount]{Status: true, Message: "", Data: &messageCounts}
}

func CreateMessage(message *models.Message, accessToken string, tokenPayload *dtos.TokenPayload) dtos.Result {
	log.Printf("[MessageService] User %d is sending a message to %d\n", tokenPayload.Id, message.ToId)

	// From id is used from token for security
	message.FromId = tokenPayload.Id
	message.ReadAt = time.Now()
	message.MessageType = "TEXT"
	message.MessageStatus = 0
	
	// Fist save entity to db to get id
	database.Connection.Create(&message)
	
	// Publish to queue if user is connected
	userOnlineStatusResult := GetUserOnlineStatus(message.ToId, accessToken)
	if  userOnlineStatusResult.Status && userOnlineStatusResult.Data.OnlineStatus == enums.USER_ONLINE {
		byteBuffer := new(bytes.Buffer)
		json.NewEncoder(byteBuffer).Encode(message)
		rabbitmq.PublishMessage(byteBuffer.Bytes(), rabbitmq.GetQueueNameForUser(message.ToId))

		return dtos.Result{Status: true, Message: "Message created and queued."}
	}

	return dtos.Result{Status: true, Message: "Message created."}
}
