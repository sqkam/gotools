package tools

import "encoding/json"

// 类型一致
// interface
// map
func JsonToAny[T any](data []byte) (T, error) {
	t := new(T)
	err := json.Unmarshal(data, t)
	return *t, err
}
func AnyToJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)

}
func AnyToAny[T, U any](u U, additionFunc ...func(a T, b U) T) (T, error) {
	kk := new(T)
	toJson, err := AnyToJson(u)
	if err != nil {
		return *kk, err
	}
	t, err := JsonToAny[T](toJson)
	if err != nil {
		return t, err
	}
	for _, f := range additionFunc {
		t = f(t, u)
	}
	return t, nil

}
