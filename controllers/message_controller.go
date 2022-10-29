package controllers

import (
	"encoding/json"
	"net/http"
	"zup-message-service/data/dtos"
	"zup-message-service/data/enums"
	"zup-message-service/data/models"
	"zup-message-service/services"

	"strconv"

	"github.com/gorilla/mux"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)

	token := r.Context().Value(enums.TOKEN).(string)
	tokenPayload := r.Context().Value(enums.TOKEN_PAYLOAD).(*dtos.TokenPayload)

	services.CreateMessage(&message, token, tokenPayload)
	json.NewEncoder(w).Encode(message)
}

func GetConversation(w http.ResponseWriter, r *http.Request) {
	toId, _ := strconv.ParseUint(mux.Vars(r)["toId"], 0, 8)
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	var pagination dtos.Pagination
	pagination.Size = size
	pagination.Page = page

	tokenPayload := r.Context().Value(enums.TOKEN_PAYLOAD).(*dtos.TokenPayload)

	messages := services.GetConversation(toId, pagination, tokenPayload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func SetMessageAsRead(w http.ResponseWriter, r *http.Request) {
	messageId, _ := strconv.ParseUint(mux.Vars(r)["id"], 0, 8)

	tokenPayload := r.Context().Value(enums.TOKEN_PAYLOAD).(*dtos.TokenPayload)

	result := services.SetMessageAsRead(messageId, tokenPayload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func SetMessagesAsRead(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(r.URL.Query().Get("userId"), 0, 8)

	tokenPayload := r.Context().Value(enums.TOKEN_PAYLOAD).(*dtos.TokenPayload)

	result := services.SetMessagesAsRead(userId, tokenPayload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
