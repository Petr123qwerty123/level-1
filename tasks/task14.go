package tasks

import (
	"reflect"
)

func Task14(x interface{}) string {
	t := reflect.TypeOf(x)

	switch t.Kind() {
	case reflect.Int:
		return "Type: int"
	case reflect.String:
		return "Type: string"
	case reflect.Bool:
		return "Type: bool"
	case reflect.Chan:
		return "Type: channel"
	default:
		return "Type: unknown"
	}
}
