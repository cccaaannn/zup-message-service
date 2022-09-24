package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterMessageRoutes(router *mux.Router) {
	router.HandleFunc("/send", controllers.SendMessage).Methods("POST")
	router.HandleFunc("/SetMessageAsRead/{id}", controllers.SetMessageAsRead).Methods("PUT")
	router.HandleFunc("/GetAllMessages/{toId}", controllers.GetAllMessages).Methods("GET")
}
