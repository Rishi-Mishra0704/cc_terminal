package utils

import (
	"cc_terminal/controllers"
	"log"
	"net/http"
	// Import gorilla/websocket package
)

// Then in your StartRouter function, use GinHandleWebSocket instead
func StartRouter() {
	http.HandleFunc("/ws", controllers.HandleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
