package utils

import (
	"cc_terminal/controllers"
	"fmt"

	"github.com/gin-gonic/gin" // Import gorilla/websocket package
	"gorm.io/gorm"
)

func GinHandleWebSocket(c *gin.Context) {
	// Upgrade the HTTP connection to a WebSocket connection
	w := c.Writer
	r := c.Request
	conn, err := controllers.Upgrader.Upgrade(w, r, nil) // Qualify upgrader with package name
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
		commandOutput := controllers.ExecuteCommand(string(message)) // Qualify executeCommand with package name

		// Send the command output back to the client
		err = conn.WriteMessage(messageType, []byte(commandOutput))
		if err != nil {
			fmt.Println("Error writing message to WebSocket:", err)
			break
		}
	}
}

// Then in your StartRouter function, use GinHandleWebSocket instead
func StartRouter(db *gorm.DB) {
	router := gin.Default()

	// Register the route handlers
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	router.GET("/ws", GinHandleWebSocket) // Use GinHandleWebSocket here

	// Start the router
	router.Run(":8080")
}
