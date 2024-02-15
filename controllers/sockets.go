package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	for {
		// Read command from client
		_, command, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading command:", err)
			break
		}

		// Execute command
		output, err := executeCommand(string(command))
		if err != nil {
			log.Println("Error executing command:", err)
			output = []byte(fmt.Sprintf("Error executing command: %s", err))
		}

		// Send output back to client
		if err := conn.WriteMessage(websocket.TextMessage, output); err != nil {
			log.Println("Error writing output:", err)
			break
		}
	}
}

func executeCommand(command string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
