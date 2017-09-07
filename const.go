package main

import (
	"fmt"
	"unsafe"
)

//常量用作枚举
const (
	UNKNOWN = 0
	MALE    = 1
	FEMALE  = 2
)

//常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)

//特殊常量：可以用作枚举
//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
const (
	e = iota //0
	f = iota //1
	g = iota //2
)

const (
	h = iota //重置为0
	i = iota //1
	j = iota //2
)

//简写形式
const (
	k = iota //重置为0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	l        //1
	m        //2
)

//iota用作性别枚举
const (
	GENDER_UNKNOWN = iota
	GENDER_MALE
	GENDER_FAMALE
)

func main() {
	const c1 uint32 = 30
	const c2 = 31
	const c3, c4 = false, "a" //常量不是必须使用的
	//c1 = 10常量不能修改，赋值报错

	fmt.Println("const", c1)
	fmt.Println("consts", c2, c3)
	fmt.Println("global consts", a, b, c)

	fmt.Println("global iota1", e, f, g)
	fmt.Println("global iota2", h, i, j)
	fmt.Println("global gender", GENDER_UNKNOWN, GENDER_MALE, GENDER_FAMALE)
}
