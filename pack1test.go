package main

import (
	"./pack"
	//. "./pack" //使用.来做为包的别名，可以不通过包名来使用其中的变量或函数，Pack1Calc()
	"fmt"
)

//import . "./pack" //使用.来做为包的别名，可以不通过包名来使用其中的变量或函数，Pack1Calc()

//主程序利用的包必须在主程序编写之前被编译
func main() {
	var a int32 = 10
	var b float32 = 31.121
	//fmt.Println(Pack1Calc(a, b)) import . "./pack"
	fmt.Println(pack1.Pack1Calc(a, b))
}
