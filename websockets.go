package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define a websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define a WebSocket handler function
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Handle incoming messages on the WebSocket connection
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// Handle the message based on its type
		switch messageType {
		case websocket.TextMessage:
			// Handle text messages
			log.Printf("Received message: %s", string(message))
		case websocket.BinaryMessage:
			// Handle binary messages (e.g., video data)
			log.Printf("Received binary message: %v bytes", len(message))
		default:
			log.Printf("Received unknown message type: %v", messageType)
		}
	}
}

func main() {
	// Define a WebSocket endpoint
	http.HandleFunc("/ws", handleWebSocket)

	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
