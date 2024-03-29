package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterMessageRoutes(router *mux.Router) {
	router.HandleFunc("/send", controllers.SendMessage).Methods("POST")
	router.HandleFunc("/read/{id}", controllers.SetMessageAsRead).Methods("PUT")
	router.HandleFunc("/read-all", controllers.SetMessagesAsRead).Methods("PUT")
	router.HandleFunc("/conversation/{toId}", controllers.GetConversation).Methods("GET")
	router.HandleFunc("/unread", controllers.GetUnreadMessageCount).Methods("GET")
}
