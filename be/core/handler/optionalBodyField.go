package handler

import (
	"bytes"
	"encoding/json"
)

type IOptionalField interface {
	GetIsSet() bool
	GetIsNull() bool
	GetValue() any
}

type OptionalBodyField[T any] struct {
	Value  T
	IsSet  bool
	IsNull bool
}

func (f *OptionalBodyField[T]) GetIsSet() bool  { return f.IsSet }
func (f *OptionalBodyField[T]) GetIsNull() bool { return f.IsNull }
func (f *OptionalBodyField[T]) GetValue() T     { return f.Value }

func (f *OptionalBodyField[T]) UnmarshalJSON(data []byte) error {
	f.IsSet = true

	if bytes.Equal(data, []byte("null")) {
		f.IsNull = true
		return nil
	}

	return json.Unmarshal(data, &f.Value)
}
