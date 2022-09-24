package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"zup-message-service/data/enums"
	"zup-message-service/services"
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

func WS(w http.ResponseWriter, r *http.Request) {

	// Raw js websocket api can not send authorization header, token is provided via url
	token := mux.Vars(r)["token"]

	tokenPayload := services.IsAuthorized(token)
	if !tokenPayload.Status {
		log.Println(tokenPayload.Message)
		return
	}
	userId := tokenPayload.Data.Id

	// ws connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	// rabbitmq channel
	channel, err := rabbitmq.Connection.Channel()
	if err != nil {
		log.Println(err)
	}

	defer channel.Close()

	// Get consumer
	messages, err := rabbitmq.GetConsumer(channel, strconv.FormatInt(int64(userId), 10))

	// Read from queue
	if err == nil {

		// make user online
		services.SetUserOnlineStatus(userId, enums.USER_ONLINE, token)

		// dc checker
		go closeOnDisconnect(conn, channel, userId, token)

		for message := range messages {
			json_message, _ := json.Marshal(string(message.Body))

			log.Println("ALIVE")

			err := conn.WriteMessage(1, json_message)

			if err != nil {
				log.Println(err)
			}
		}
	}

}

func closeOnDisconnect(conn *websocket.Conn, channel *amqp.Channel, userId uint64, token string) {
	for {
		time.Sleep(time.Second)
		_, _, err := conn.ReadMessage()
		if err != nil {
			channel.Close()
			conn.Close()
			services.SetUserOnlineStatus(userId, enums.USER_OFFLINE, token)
			break
		}
	}
}
