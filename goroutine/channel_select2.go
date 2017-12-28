package main

import (
	"fmt"
	"os"
)

func main() {
	c1, c2 := make(chan int, 3), make(chan int)

	//sub goroutine 消费数据
	go func() {
		d, ok, s := 0, false, ""
		for {
			select { //随机选择可用的channel，读取数据
			case d, ok = <-c1:
				s = "c1"
			case d, ok = <-c2:
				s = "c2"
			default:
				fmt.Println("no active channel to read")
			}
			if ok {
				fmt.Println(s, d)
			} else {
				fmt.Println(s, "is closed")
				os.Exit(0)
			}
		}
	}()

	//main goroutine 生产数据
	for i := 0; i < 5; i++ { //在循环中使用 select default case 需要小心，避免形成洪水。
		select { //随机选择可用的channel，发送数据
		case c1 <- i:
		case c2 <- i:
			//default:
			//	fmt.Println("no active channel to write")
		}
	}
	//close(c1)

	//select {} //没有可用 channel，阻塞 main goroutine。
}
