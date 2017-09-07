package main

import "fmt"

//因式分解关键字：用于声明全局变量，允许仅申明，不使用
var (
	g1 int16  = 1
	g2 uint32 = 10
	g3        = 1
	g4 bool
)

func main() {

	fmt.Println(g1, g2, g3, g4)

	//使用var定义变量，指定类型
	var a int32
	var s string

	//同一类型的，申明在一行
	var i1, i2, i3 int
	//同一行进行同时赋值，并行赋值也被用于一个函数有多个返回值时
	i1, i2, i3 = 1, 2, 3
	//交换两个变量
	i1, i2 = i2, i1
	//空白标识符，用于抛弃
	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	_, i1 = -1, 100
	fmt.Println(i1, i2, i3)

	//类型推断
	var b1 = false
	fmt.Println(b1)

	a = 100
	s = "I am string"
	fmt.Println(a, s)

	//推荐的简短形式：申明并赋值，只能用在函数体内，前面已申明的变量名称，不能再重复申明
	b := 200
	//b := 200//已申明，不能再次声明
	ss := "I am string2"

	fmt.Println(b, ss)
}
