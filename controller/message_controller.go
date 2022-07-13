package controller

import (
	"encoding/json"
	"net/http"
	"zup-message-service/database"
	"zup-message-service/model"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []model.Message
	database.Connection.Find(&messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
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
