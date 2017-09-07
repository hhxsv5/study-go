package main

import (
	"errors"
	"fmt"
)

//error类型是一个接口类型
//type error interface {
//	Error() string
//}

//自定义除数错误

type DivideError struct {
	a, b float64
}

func (err DivideError) Error() string {
	return fmt.Sprintf("%f和%f不能除", err.a, err.b)
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New(fmt.Sprintf("%f和%f不能除", a, b))
	}
	return a / b, nil
}

func divideNew(a float64, b float64) (result float64, err string) {
	if b == 0 {
		dErr := DivideError{a, b}
		err = dErr.Error()
		return
	}
	//return a / b, ""
	result, err = a/b, ""
	return
}

func main() {
	a, b := 10.0, 0.0
	result, err1 := divide(a, b)
	//fmt.Println(result, err1)
	if err1 == nil {
		fmt.Println(a, "/", b, "=", result)
	} else {
		fmt.Println("Error: ", err1)
	}

	result, err2 := divideNew(a, b)
	//fmt.Println(result, err2)
	if err2 == "" {
		fmt.Println(a, "/", b, "=", result)
	} else {
		fmt.Println("Error: ", err2)
	}

}
