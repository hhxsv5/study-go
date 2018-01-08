package main

import "fmt"

func main() {
	//byte 字节型，实际是uint8，用来强调数据是raw data，而不是数字
	s := "go编程"
	b := []byte(s) //string转[]byte
	fmt.Println(b) //每个元素是uint8的数字

	fmt.Println(string(b)) //[]byte转string
}
