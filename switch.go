package main

import (
	"fmt"
)

func main() {

	x := 3

	//switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break
	switch x {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("case default")
	}

	var y interface{}

	switch i := y.(type) {
	case nil:
		fmt.Println("%T", i)
	case int:
		fmt.Println("int类型")
	case float32:
		fmt.Println("float32类型")
	case float64:
		fmt.Println("float64类型")
	case func(int) float64:
		fmt.Println("func(int)返回值类型为float64")
	case bool, string:
		fmt.Println("bool或string类型")
	default:
		fmt.Println("未知类型")
	}
}
