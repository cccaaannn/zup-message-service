package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"zup-message-service/rabbitmq"

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

	// TODO get userId from token
	userId, err := strconv.ParseUint(mux.Vars(r)["user"], 0, 0)
	log.Println("Client Connected")

	// Get consumer
	messages, err := rabbitmq.GetConsumer(channel, strconv.FormatUint(userId, 10))

	// Read from queue
	if err == nil {
		go closeOnDisconnect(conn, channel)

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

func closeOnDisconnect(conn *websocket.Conn, channel *amqp.Channel) {
	for {
		time.Sleep(time.Second)
		_, _, err := conn.ReadMessage()
		if err != nil {
			channel.Close()
			conn.Close()
			break
		}
	}
}
