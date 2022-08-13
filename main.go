package main

import (
	"fmt"
	"log"
	"net/http"
	"zup-message-service/configs"
	"zup-message-service/database"
	"zup-message-service/middlewares"
	"zup-message-service/rabbitmq"
	"zup-message-service/routes"

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
	userRouter := router.PathPrefix("/api/v1/messages").Subrouter()
	wsRouter := router.PathPrefix("/api/v1/ws").Subrouter()

	routes.RegisterMessageRoutes(messageRouter)
	routes.RegisterUserOnlineStatusRoutes(userRouter)
	routes.RegisterWsRoutes(wsRouter)

	messageRouter.Use(middlewares.AuthenticationMiddleware)
	userRouter.Use(middlewares.AuthenticationMiddleware)

	// cors
	handler := cors.AllowAll().Handler(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", configs.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configs.AppConfig.Port), handler))

}

// import (
// 	"fmt"
// 	"zup-message-service/services"
// )

// func main() {
// 	result := services.IsUserAuthorized()

// 	if result.Status {
// 		fmt.Printf("%s\n", result.Data.Username)
// 	} else {
// 		fmt.Printf("%+v\n", result)
// 	}

// }
