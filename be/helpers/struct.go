package helpers

import "reflect"

func StructToQueryFilters(data interface{}) map[string][]string {
	result := make(map[string][]string)

	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		tag := structField.Tag.Get("form")

		if tag == "" || field.Len() == 0 {
			continue
		}

		if strSlice, ok := field.Interface().([]string); ok {
			result[tag] = strSlice
		}
	}

	return result
}
