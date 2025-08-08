package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type message struct {
	User  string `json:"user"`
	Value string `json:"value"`
}

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
	go sendWelcomeMessage(ws)
	go listenOnCliInput(ws)

	listenOnClient(ws)
}

func sendWelcomeMessage(ws *websocket.Conn) {
	welcomeMessage := message{"SERVER", "Hello there 👋"}

	err := ws.WriteJSON(welcomeMessage)
	if err != nil {
		log.Println(err)
	}
}

func listenOnCliInput(ws *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		input, _ := reader.ReadString('\n')
		var message = message{"SERVER", input}

		err := ws.WriteJSON(message)
		if err != nil {
			log.Println(err)
		}
	}
}

func listenOnClient(ws *websocket.Conn) {
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := ws.WriteMessage(messageType, p); err != nil {
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
