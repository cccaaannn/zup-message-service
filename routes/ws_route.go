package routes

import (
	"zup-message-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterWsRoutes(router *mux.Router) {
	router.HandleFunc("/{user}", controllers.WS)
}
