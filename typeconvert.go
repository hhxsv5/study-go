package main

import "fmt"

//类型转换
func main() {
	var a int = 10
	var b float32 = 3
	var sum float64

	//只有ab连个操作数类型相同才能做运算
	//sum = a / b

	sum = float64(a) / float64(b)
	fmt.Println(a, "/", b, "=", sum)
}
