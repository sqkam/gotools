package gotools

import "strconv"

var FuncString2Float64 = func(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
