package core

import (
	"bytes"
	"encoding/json"
)

type OptionalBodyField[T any] struct {
	Value  T
	IsSet  bool
	IsNull bool
}

func (o *OptionalBodyField[T]) UnmarshalJSON(data []byte) error {
	o.IsSet = true

	if bytes.Equal(data, []byte("null")) {
		o.IsNull = true
		return nil
	}

	return json.Unmarshal(data, &o.Value)
}
