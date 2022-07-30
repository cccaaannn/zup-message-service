package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserOnlineStatusRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/userStatus/getStatus/{userId}", controllers.GetLastUserLogin).Methods("GET")
	router.HandleFunc("/api/v1/userStatus/update/{userId}/{status}", controllers.UpdateLastUserLogin).Methods("PUT")
}
