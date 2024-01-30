package gql

import (
	"fmt"
	"reflect"
)

type (
	ID      string
	Int     int32
	Float   float64
	String  string
	Boolean bool
)

func NewID(v any) *ID {
	id := ToID(v)
	return &id
}

func ToID(v any) ID {
	var str string
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = fmt.Sprintf("%d", v)
		if str == "0" {
			str = ""
		}
	case reflect.String:
		str = v.(string)
	}
	return ID(str)
}
