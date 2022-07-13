package main

import (
	"fmt"
	"log"
	"net/http"
	"zup-message-service/config"
	"zup-message-service/controller"
	"zup-message-service/database"

	"github.com/gorilla/mux"
)

func main() {

	config.LoadConfig()

	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterMessageRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}

func RegisterMessageRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/messages", controller.GetMessages).Methods("GET")
	// router.HandleFunc("/api/messages/{id}", controller.GetMessageById).Methods("GET")
	router.HandleFunc("/api/v1/messages", controller.CreateMessage).Methods("POST")
}
