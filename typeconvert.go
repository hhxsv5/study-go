package main

import (
	"fmt"
	"strconv"
)

//类型转换
func main() {
	var a int = 10
	var b float32 = 3
	var sum float64
	var str string = "12"
	var c64 int64 = 1233

	//只有ab连个操作数类型相同才能做运算
	//sum = a / b

	sum = float64(a) / float64(b)
	fmt.Println(a, "/", b, "=", sum)

	//int to string
	fmt.Println(strconv.Itoa(a))

	//int64 to string
	fmt.Println(strconv.FormatInt(c64, 10))

	//string to int
	fmt.Println(strconv.Atoi(str))

	//string to int64
	fmt.Println(strconv.ParseFloat(str, 64))
}
