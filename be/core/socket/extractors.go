package socket

import "encoding/json"

func GetMessagePayload[T any](messagePayload json.RawMessage) (T, error) {
	var payload T

	if err := json.Unmarshal(messagePayload, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
