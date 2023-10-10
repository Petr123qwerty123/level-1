package tasks

import (
	"reflect"
)

// Task14 принимает на вход переменную любого типа, на выходе возвращает тип переменной в виде строки
func Task14(x any) string {
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
