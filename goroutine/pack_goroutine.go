package main

import (
	"fmt"
	"io/ioutil"
)

//用简单工厂模式打包并发任务和 channel
func readIo() chan string {
	c := make(chan string)

	go func() {
		d, e := ioutil.ReadFile("io.data")
		if e != nil {
			panic(e)
		}
		c <- string(d)
	}()

	return c
}

func main() {
	c := readIo()
	fmt.Println(<-c)
}
