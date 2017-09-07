package main

import "fmt"

func main() {
	var x int = 10
	//申明指针，初始化为空指针，值为nil
	var px *int
	//取x的地址
	px = &x
	fmt.Println(px)
	//获取指针指向地址的值
	fmt.Println(*px)

	//修改指针指向地址的值
	*px = 20
	fmt.Println(*px, x)

	if px != nil {
		fmt.Println("非空指针")
	}

	//指向指针的指针，多级指针
	var ppx **int
	ppx = &px
	fmt.Println(ppx, *ppx, **ppx)

	var array = [4]int32{1, 2, 3, 4}
	var pa = &array
	fmt.Println(pa)
	for i := 0; i < len(pa); i++ {
		fmt.Println(pa[i])
	}

	//指针参数
	incr(&x)
	fmt.Println(x)
}

//指针作为函数参数
func incr(p *int) {
	*p += 1
}
