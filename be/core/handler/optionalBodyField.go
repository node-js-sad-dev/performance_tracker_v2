package handler

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
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
func (f *OptionalBodyField[T]) GetValue() any   { return f.Value }

func (f *OptionalBodyField[T]) UnmarshalJSON(data []byte) error {
	f.IsSet = true

	if bytes.Equal(data, []byte("null")) {
		f.IsNull = true
		return nil
	}

	return json.Unmarshal(data, &f.Value)
}

func (f *OptionalBodyField[T]) UnmarshalText(text []byte) error {
	f.IsSet = true

	strText := string(text)

	if strText == "" || strText == "null" {
		f.IsNull = true
		return nil
	}
	f.IsNull = false

	var val T
	ptr := any(&val)

	if unmarshaler, ok := ptr.(encoding.TextUnmarshaler); ok {
		if err := unmarshaler.UnmarshalText(text); err != nil {
			return err
		}
		f.Value = val
		return nil
	}

	rv := reflect.ValueOf(ptr).Elem()

	switch rv.Kind() {
	case reflect.String:
		rv.SetString(strText)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parsed, err := strconv.ParseInt(strText, 10, rv.Type().Bits())
		if err != nil {
			return fmt.Errorf("failed to parse %q as int: %w", strText, err)
		}
		rv.SetInt(parsed)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		parsed, err := strconv.ParseUint(strText, 10, rv.Type().Bits())
		if err != nil {
			return fmt.Errorf("failed to parse %q as uint: %w", strText, err)
		}
		rv.SetUint(parsed)

	case reflect.Float32, reflect.Float64:
		parsed, err := strconv.ParseFloat(strText, rv.Type().Bits())
		if err != nil {
			return fmt.Errorf("failed to parse %q as float: %w", strText, err)
		}
		rv.SetFloat(parsed)

	case reflect.Bool:
		parsed, err := strconv.ParseBool(strText)
		if err != nil {
			return fmt.Errorf("failed to parse %q as bool: %w", strText, err)
		}
		rv.SetBool(parsed)

	default:
		return fmt.Errorf("type %T does not implement encoding.TextUnmarshaler and is not a supported primitive type", val)
	}

	f.Value = val
	return nil
}
