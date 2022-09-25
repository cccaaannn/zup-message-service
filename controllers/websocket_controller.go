package controllers

import (
	"net/http"
	"zup-message-service/services"
)

func HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	services.HandleWebsocketConnection(w, r)
}
