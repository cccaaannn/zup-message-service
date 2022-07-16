package controller

import (
	"encoding/json"
	"net/http"
	"zup-message-service/database"
	"zup-message-service/model"
	"zup-message-service/service"

	"strconv"

	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []model.Message
	database.Connection.Find(&messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var messages []model.Message
	fromId, _ := strconv.ParseUint(mux.Vars(r)["fromId"], 0, 8)
	toId, _ := strconv.ParseUint(mux.Vars(r)["toId"], 0, 8)

	messages = service.GetAllMessages(fromId, toId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func SetMessageAsRead(w http.ResponseWriter, r *http.Request) {
	messageId, _ := strconv.ParseUint(mux.Vars(r)["id"], 0, 8)
	result := service.SetMessageAsRead(messageId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// func GetMessageById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfProductExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Product
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message model.Message
	json.NewDecoder(r.Body).Decode(&message)
	database.Connection.Create(&message)
	json.NewEncoder(w).Encode(message)
}
