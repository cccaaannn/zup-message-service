package services

import (
	"fmt"
	"log"
    "encoding/json"

	"zup-message-service/configs"
	"zup-message-service/data/dtos"
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func getUserServiceBasePath() string {
	return fmt.Sprintf("%s%s", configs.AppConfig.UserServiceUrl, configs.AppConfig.UserServiceApiPathPrefix)
}

func IsAuthorized(userToken string) dtos.DataResult[dtos.TokenPayload] {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/authorization/%s", userToken)
	log.Printf("Requesting user service, URL: %s\n", requestURL)
	res := dtos.DataResult[dtos.TokenPayload]{Status:false, Message: "Error", Data: nil}
	
	resp, err := client.R().EnableTrace().Get(requestURL)

	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(resp.Body(), &res)
    if err != nil {
        fmt.Println(err)
    }

	log.Printf("Response from user service: %+v\n\n", res)
	return res
}

func SetUserOnlineStatus(userId uint64, newStatus string, accessToken string) dtos.Result {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/users/%d/online-status/%s", userId, newStatus)
	log.Printf("Requesting user service, URL: %s\n", requestURL)
	res := dtos.Result{Status:false, Message: "Error"}

	resp, err := client.R().
		SetAuthToken(accessToken).
      	Put(requestURL)

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Response from user service: %+v\n\n", res)
	return res
}

func GetUserOnlineStatus(userId uint64, accessToken string) dtos.DataResult[dtos.UserOnlineStatus] {
	requestURL := fmt.Sprintf(getUserServiceBasePath() + "/users/%d/online-status", userId)
	log.Printf("Requesting user service, URL: %s\n", requestURL)
	res := dtos.DataResult[dtos.UserOnlineStatus]{Status:false, Message: "Error", Data: nil}

	resp, err := client.R().
		SetAuthToken(accessToken).
      	Get(requestURL)

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Response from user service: %+v\n\n", res)
	return res
}
