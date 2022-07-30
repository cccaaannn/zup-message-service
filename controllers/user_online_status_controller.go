package controllers

import (
	"encoding/json"
	"net/http"
	"zup-message-service/services"

	"strconv"

	"github.com/gorilla/mux"
)

func UpdateLastUserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	status, _ := strconv.Atoi(mux.Vars(r)["status"])

	res := services.UpdateLastUserLogin(userId, status)

	json.NewEncoder(w).Encode(res)
}

func GetLastUserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	userOnlineStatus := services.GetLastUserLogin(userId)

	json.NewEncoder(w).Encode(userOnlineStatus)
}
