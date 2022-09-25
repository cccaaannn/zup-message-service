package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterWebSocketRoutes(router *mux.Router) {
	router.HandleFunc("/{token}", controllers.HandleWebsocketConnection)
}
