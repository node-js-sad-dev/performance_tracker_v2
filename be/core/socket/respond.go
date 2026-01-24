package socket

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// for now commiting crime by adding package i don't plan to use, will remove it from here later when i will have basic project structure i am happy with

func Respond(conn *websocket.Conn, success bool, payload interface{}, err error) {
	msg := gin.H{
		"success": success,
	}
	if success {
		msg["data"] = payload
	} else if err != nil {
		msg["error"] = err.Error()
	} else {
		msg["error"] = "unknown error"
	}

	if writeErr := conn.WriteJSON(msg); writeErr != nil {
		log.Printf("Failed to send message over WebSocket: %v", writeErr)
	}
}
