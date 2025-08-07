package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there 👋")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// TODO: refactore later to a valid cors config!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected ✅")

	err = ws.WriteMessage(1, []byte("[server message] Hey there client!"))
	if err != nil {
		log.Println(err)
	}

	listenOnClient(ws)
}

func listenOnClient(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
