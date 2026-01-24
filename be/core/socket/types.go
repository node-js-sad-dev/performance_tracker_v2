package socket

import (
	"encoding/json"
)

type ClientMessage struct {
	Type    string          `json:"type"`    // e.g. "auth", "chat", "ping"
	Payload json.RawMessage `json:"payload"` // raw payload to decode based on type
}

type AuthSocketMessagePayload struct {
	Token string `json:"token"`
}
