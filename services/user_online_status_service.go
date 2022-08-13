package services

import (
	"time"
	"zup-message-service/database"
	"zup-message-service/models"
)

func UpdateLastUserLogin(userId uint64, status int) bool {
	var userOnlineStatus models.UserOnlineStatus
	userOnlineStatus.UserId = userId
	userOnlineStatus.IsOnline = status
	userOnlineStatus.LastOnline = time.Now()

	database.Connection.Save(&userOnlineStatus)
	return true
}

func GetLastUserLogin(userId int) models.UserOnlineStatus {
	var userOnlineStatus models.UserOnlineStatus
	database.Connection.First(&userOnlineStatus, userId)
	return userOnlineStatus
}
