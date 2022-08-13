package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"zup-message-service/rabbitmq"
	"zup-message-service/services"

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

	user := services.IsUserAuthorized(mux.Vars(r)["user"])
	if !user.Status {
		log.Println(user.Message)
		return
	}
	userId := user.Data.Id

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
		services.UpdateLastUserLogin(userId, 1)

		// dc checker
		go closeOnDisconnect(conn, channel, userId)

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

func closeOnDisconnect(conn *websocket.Conn, channel *amqp.Channel, userId uint64) {
	for {
		time.Sleep(time.Second)
		_, _, err := conn.ReadMessage()
		if err != nil {
			channel.Close()
			conn.Close()
			services.UpdateLastUserLogin(userId, 0)
			break
		}
	}
}
