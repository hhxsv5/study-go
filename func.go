package main

import (
	"fmt"
	"math"
)

//函数声明告诉了编译器函数的名称，返回类型，和参数。
//单参数，单返回值
func increment(a int) int {
	a++
	return a
}

//多参数，单返回值
func sum(a int, b int) int {
	return a + b
}

//多参数，多返回值，不指定返回值名称
func increments1(a int, b int) (int, int) {
	a++
	b++
	return a, b
}

//多参数，多返回值，指定返回值名称
func increments2(a int, b int) (a1 int, b1 int) {
	a++
	b++
	return a, b
}

//闭包：Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func increment3(start int) func() int {
	i := start
	return func() int {
		i++
		return i
	}
}

//定义结构圆
type Circle struct {
	radius float64
}

//计算圆面积
//方法：一个包含了接受者(c Circle)的函数
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

	a, b := 200, -100
	//函数作为值
	fmt.Println(increment(a), increment(b))
	fmt.Println(sum(a, b))

	aa1, bb1 := increments1(a, b)
	fmt.Println(aa1, bb1)

	aa2, bb2 := increments2(a, b)
	fmt.Println(increments2(aa2, bb2))

	//函数作为闭包
	increment := increment3(0)
	fmt.Println(increment()) //1
	fmt.Println(increment()) //2
	fmt.Println(increment()) //3

	//函数作为方法
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆面积为", c1.getArea())
}

//函数内定义的变量称为局部变量：作用域只在函数体内，参数和返回值变量也是局部变量。
//函数外定义的变量称为全局变量：全局变量可以在整个包甚至外部包（被导出后）使用
//函数定义中的变量称为形式参数：作用域和局部变量一样
//Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
//int float32默认0，pointer默认nil
