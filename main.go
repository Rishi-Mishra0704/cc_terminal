package main

import (
	"cc_terminal/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", controllers.HandleWebSocket)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
