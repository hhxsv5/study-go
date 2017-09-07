package main

import "fmt"

func main() {
	a, b := 11, 2

	c := a + b
	fmt.Println(a, "+", b, "=", c)

	c = a * b
	fmt.Println(a, "*", b, "=", c)

	c = a % b
	fmt.Println(a, "%", b, "=", c)

	if a >= b {
		fmt.Println(a, ">=", b)
	}

	if a != b {
		fmt.Println(a, "!=", b)
	}

	d := a
	if a == d {
		fmt.Println(a, "==", d)
	}

	if a > 0 && d > 0 {
		fmt.Println(a, ">0&&", b, ">0=true")
	}

	d += 10
	a -= 20

	//取变量存储地址
	ptr := &a
	fmt.Println("a ptr", ptr, *ptr)
}
