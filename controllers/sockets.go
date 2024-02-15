package controllers

import (
	"cc_terminal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message from WebSocket:", err)
			break
		}

		// Split the message into command and arguments
		parts := strings.Fields(string(message))
		command := parts[0]
		args := parts[1:]

		var output string
		switch command {
		case "ls":
			output = utils.ListFiles()
		case "mkdir":
			output = utils.CreateDirectory(args)
		case "cd":
			output = utils.ChangeDirectory(args)
		case "echo":
			output = utils.Echo(args)
		case "pwd":
			output = utils.Pwd()
		default:
			output = utils.ExecuteCommand(command, args)
		}

		err = conn.WriteMessage(messageType, []byte(output))
		if err != nil {
			fmt.Println("Error writing message to WebSocket:", err)
			break
		}
	}
}
