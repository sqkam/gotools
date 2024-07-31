package tools

import "reflect"

func NewAny[T any]() (*T, error) {
	str := "{}"
	t := new(T)
	val := indirect(reflect.ValueOf(t))
	if val.Kind() == reflect.Slice {
		str = "[]"
	}
	toAny, err := JsonToAny[T]([]byte(str))
	return &toAny, err
}
