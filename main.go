package main

import (
	"fmt"
	"log"
	"net/http"
	"zup-message-service/configs"
	"zup-message-service/database"
	"zup-message-service/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	configs.LoadConfig()

	// Db connection
	database.Connect(configs.AppConfig.PostgresqlConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	routes.RegisterMessageRoutes(router)

	// cors
	handler := cors.AllowAll().Handler(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", configs.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configs.AppConfig.Port), handler))

}
