package main

import (
	"fmt"
	"log"
	"net/http"
	"zup-message-service/configs"
	"zup-message-service/middlewares"
	"zup-message-service/routes"
	"zup-message-service/utils/database"
	"zup-message-service/utils/rabbitmq"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	configs.LoadConfig()

	// Db connection
	database.Connect(configs.AppConfig.PostgresqlConnectionString)
	database.Migrate()

	// RabbitMQ connection
	rabbitmq.Connect(configs.AppConfig.RabbitmqConnectionString)
	defer rabbitmq.Disconnect()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	messageRouter := router.PathPrefix("/api/v1/messages").Subrouter()
	webSocketRouter := router.PathPrefix("/api/v1/ws").Subrouter()

	routes.RegisterMessageRoutes(messageRouter)
	routes.RegisterWebSocketRoutes(webSocketRouter)

	messageRouter.Use(middlewares.AuthorizationMiddleware)

	// cors
	handler := cors.AllowAll().Handler(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", configs.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configs.AppConfig.Port), handler))

}
