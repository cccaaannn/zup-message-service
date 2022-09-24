package services

import (
	"fmt"
	"log"

	"zup-message-service/utils"
	"zup-message-service/configs"
	"zup-message-service/data/dtos"
)


var userServiceApiPath = "api/v1"

func getUserServiceBasePath() string {
	return fmt.Sprintf("%s:%s/%s", configs.AppConfig.UserServiceUrl, configs.AppConfig.UserServicePort, userServiceApiPath)
}

func IsAuthorized(userToken string) dtos.DataResult[dtos.TokenPayload] {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/authorization/%s", userToken)
	log.Println(requestURL)

	return utils.GetApiDataResult[dtos.TokenPayload](requestURL)
}

func SetUserOnlineStatus(userId uint64, newStatus string, accessToken string) dtos.Result {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/user/%d/online-status/%s", userId, newStatus)
	log.Println(requestURL)

	return utils.PutApiResult(requestURL, accessToken)
}

func GetUserOnlineStatus(userId uint64, accessToken string) dtos.DataResult[dtos.UserOnlineStatus] {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/user/%d/online-status", userId)
	log.Println(requestURL)

	return utils.GetApiDataResultWithToken[dtos.UserOnlineStatus](requestURL, accessToken)
}
