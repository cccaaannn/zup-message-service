package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserOnlineStatusRoutes(router *mux.Router) {
	router.HandleFunc("/getStatus/{userId}", controllers.GetLastUserLogin).Methods("GET")
	router.HandleFunc("/update/{userId}/{status}", controllers.UpdateLastUserLogin).Methods("PUT")
}
