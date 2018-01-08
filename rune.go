package main

import "fmt"

func main() {
	//Rune 是int32 的别名，字码型，代表一个Unicode码，Unicode码点
	//用于遍历字符串中的字符
	//通常，当人们讨论字符时，多数是指8 位字符。UTF-8 字符可能会有32 位，称作rune
	//在Go当中 string底层是用byte数组存的，并且是不可以改变的。
	s := "go编程"
	fmt.Println(len(s))         //输出8，一个中文占三个字节，2+3*2=8，也有个别汉字是四个字节。
	fmt.Println(len([]rune(s))) //输出4，将字符串转成rune的切片

	//用string存储unicode的话，如果有中文，按下标是无法访问得到一个中文的，因为只能取到一个byte，需要通过rune切片来按下标访问。
	fmt.Println(string(rune(s[2])), string([]rune(s)[2]))
}
