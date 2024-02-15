package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Read messages from the WebSocket connection
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message from WebSocket:", err)
			break
		}

		// Process the received message (bash command)
		commandOutput := ExecuteCommand(string(message))

		// Send the command output back to the client
		err = conn.WriteMessage(messageType, []byte(commandOutput))
		if err != nil {
			fmt.Println("Error writing message to WebSocket:", err)
			break
		}
	}
}

func ExecuteCommand(command string) string {
	// Placeholder: Execute the bash command and return the output
	// You need to implement the actual command execution logic here
	return "Command output: " + command
}
