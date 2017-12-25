package test

import "errors"

func Sum(a int32, b int32) int32 {
	return a + b
}

func Divide(a float32, b float32) float32 {
	if b == 0 {
		panic(errors.New("divide 0"))
	}
	return a / b
}
