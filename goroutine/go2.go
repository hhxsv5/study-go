package main

import (
	"fmt"
	"io/ioutil"
)

func ReadIO(ic chan string) {
	fmt.Println("Reading IO...")

	b, e := ioutil.ReadFile("/Users/dave/Documents/www/go/src/test/go/io.data")
	if e != nil {
		panic(e)
	}

	ic <- string(b) //读取到io数据
}

func main() {
	var ic = make(chan string) //创建io的chanel
	var io string

	go ReadIO(ic)

	io = <-ic //获取io数据
	fmt.Println("Read data: ", io)
}
