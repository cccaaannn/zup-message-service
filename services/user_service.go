package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"zup-message-service/configs"
	"zup-message-service/dtos"
)

var userServicePath = "api/v1/user/isUserAuthorized"

func IsUserAuthorized(userToken string) dtos.Result {
	result := dtos.Result{Status: false, Message: "", Data: nil}

	requestURL := fmt.Sprintf("%s:%s/%s/%s", configs.AppConfig.UserServiceUrl, configs.AppConfig.UserServicePort, userServicePath, userToken)
	// fmt.Printf(requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return result
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	// body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		result = dtos.Result{Status: false, Message: "", Data: nil}
	}

	// fmt.Printf("status code: %d\n", res.StatusCode)
	log.Println(result)

	return result
}
