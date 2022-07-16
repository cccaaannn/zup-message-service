package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"zup-message-service/services"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WS(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// messageId, err := strconv.ParseUint(r.Header.Get("after"), 0, 0)
	// userId, err := strconv.ParseUint(r.Header.Get("user"), 0, 0)

	messageId, err := strconv.ParseUint(mux.Vars(r)["after"], 0, 0)
	userId, err := strconv.ParseUint(mux.Vars(r)["user"], 0, 0)

	log.Println(messageId)
	log.Println(userId)

	log.Println("Client Connected")
	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws, messageId, userId)
}

func reader(conn *websocket.Conn, messageId uint64, userId uint64) {
	for {

		time.Sleep(3 * time.Second)

		messages := services.GetUnReadMessagesAfter(messageId, userId)

		json_messages, _ := json.Marshal(messages)

		err := conn.WriteMessage(1, json_messages)

		if err != nil {
			log.Println(err)
			return
		}

	}
}
