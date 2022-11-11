package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"zup-message-service/data/enums"
	"zup-message-service/utils/rabbitmq"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {

	// Raw js websocket api can not send authorization header, token is provided via url
	token := mux.Vars(r)["token"]

	tokenPayload := IsAuthorized(token)
	if !tokenPayload.Status {
		log.Println(tokenPayload.Message)
		return
	}
	userId := tokenPayload.Data.Id

	// ws connection
	ws_conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws_conn.Close()

	// rabbitmq channel
	rabbit_channel, err := rabbitmq.GetChannel()
	if err != nil {
		log.Println("[WebsocketService] RabbitMQ connection not presents, socket is not created. error:", err)
		return
	}
	defer rabbit_channel.Close()

	// Get consumer
	messages, err := rabbitmq.GetConsumer(rabbit_channel, rabbitmq.GetQueueNameForUser(userId))
	if err != nil {
		log.Println("[WebsocketService] could not get RabbitMQ consumer channel", err)
		return
	}

	// make user online
	SetUserOnlineStatus(userId, enums.USER_ONLINE, token)
	log.Printf("[WebsocketService] User %d is online.\n", userId)

	// check for dc
	go closeOnDisconnect(ws_conn, rabbit_channel, userId, token)

	for message := range messages {
		json_message, _ := json.Marshal(string(message.Body))

		log.Printf("[WebsocketService] User %d is received a message.\n", userId)

		err := ws_conn.WriteMessage(1, json_message)

		if err != nil {
			log.Println(err)
		}
	}

}

func closeOnDisconnect(ws_conn *websocket.Conn, rabbit_channel *amqp.Channel, userId uint64, token string) {
	for {
		time.Sleep(time.Second)
		_, _, err := ws_conn.ReadMessage()
		if err != nil {
			log.Printf("[WebsocketService] User %d is disconnected.\n", userId)
			rabbit_channel.Close()
			ws_conn.Close()
			SetUserOnlineStatus(userId, enums.USER_OFFLINE, token)
			break
		}
	}
}
