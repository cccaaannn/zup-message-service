package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterMessageRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/messages/send", controllers.SendMessage).Methods("POST")
	router.HandleFunc("/api/v1/messages/SetMessageAsRead/{id}", controllers.SetMessageAsRead).Methods("PUT")
	router.HandleFunc("/api/v1/messages/GetAllMessages/{fromId}/{toId}", controllers.GetAllMessages).Methods("GET")

	router.HandleFunc("/api/v1/ws/{after}/{user}", controllers.WS)
}
