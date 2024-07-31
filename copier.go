package tools

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
)

var cpyOption = copier.Option{
	IgnoreEmpty: true,
	DeepCopy:    true,
	Converters: []copier.TypeConverter{
		{
			DstType: []string{},
			SrcType: copier.String,
			Fn: func(src interface{}) (interface{}, error) {
				var des []string
				err := json.Unmarshal([]byte(src.(string)), &des)
				if err != nil {
					return nil, err
				}
				return des, err
			},
		},
		{
			DstType: copier.String,
			SrcType: []string{},
			Fn: func(src interface{}) (interface{}, error) {
				sli := src.([]string)
				jsonByte, err := json.Marshal(sli)
				if err != nil {
					return nil, errors.New("toJSON Err")
				}
				return string(jsonByte), nil
			},
		},
	},
}

func CopyAny[T any, U any](data U, additionFunc ...func(a T, b U) T) (daoModel T, err error) {
	// additionFunc is commonly used for type conversion
	tPointer, err := NewAny[T]()
	if err != nil {
		return *tPointer, err
	}
	err = copier.CopyWithOption(tPointer, data, cpyOption)
	t := *tPointer
	for _, f := range additionFunc {
		t = f(t, data)
	}
	return t, err
}
func CopyAnyErr[T any, U any](data U, additionFunc ...func(a T, b U) (T, error)) (T, error) {
	// additionFunc is commonly used for type conversion
	tPointer, err := NewAny[T]()
	if err != nil {
		return *tPointer, err
	}
	err = copier.CopyWithOption(tPointer, data, cpyOption)
	t := *tPointer
	for _, f := range additionFunc {
		t, err = f(t, data)
		if err != nil {
			return *tPointer, err
		}
	}
	return t, err
}

func CopyAnyToDest[T any, U any](t *T, data U, additionFunc ...func(a *T, b U)) error {
	// additionFunc is commonly used for type conversion
	err := copier.CopyWithOption(t, data, cpyOption)
	if err != nil {
		return err
	}

	for _, f := range additionFunc {
		f(t, data)
	}
	return nil
}

func CopySlice[T, U any](data []U, additionFunc ...func(t T, s U) T) ([]T, error) {
	// additionFunc is commonly used for type conversion
	//T is type of element of data
	s := make([]T, len(data))
	for i, v := range data {
		copyAny, err := CopyAny[T](v)
		if err != nil {
			return nil, err
		}
		s[i] = copyAny
		for _, f := range additionFunc {
			s[i] = f(s[i], v)
		}
	}
	return s, nil
}
