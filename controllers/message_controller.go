package controllers

import (
	"encoding/json"
	"net/http"
	"zup-message-service/models"
	"zup-message-service/services"
	"zup-message-service/utils"

	"strconv"

	"github.com/gorilla/mux"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)

	// TODO find a better way
	token, err := utils.GetTokenFromHeader(r)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	services.CreateMessage(&message, token)
	json.NewEncoder(w).Encode(message)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message
	fromId, _ := strconv.ParseUint(mux.Vars(r)["fromId"], 0, 8)
	toId, _ := strconv.ParseUint(mux.Vars(r)["toId"], 0, 8)

	messages = services.GetAllMessages(fromId, toId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func SetMessageAsRead(w http.ResponseWriter, r *http.Request) {
	messageId, _ := strconv.ParseUint(mux.Vars(r)["id"], 0, 8)
	result := services.SetMessageAsRead(messageId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
