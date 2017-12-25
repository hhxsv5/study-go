package main

import "fmt"

func main() {
	//通道只可使用make创建，分为有缓存通道以及无缓存通道。
	//同步模式必须有配对操作的goroutine出现，否则会一直阻塞。
	//而异步模式在缓冲区未满时或数据未被读完前不会阻塞。
	//多数情况下异步通道有助于提升性能，减少排队阻塞。
	c := make(chan int, 4) //创建一个带3个缓冲的异步通道，int后无参数为无缓存的通道，前四个元素可无阻塞写入，第五个元素写入时会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
	c <- 1                 //缓冲区未满不会阻塞
	c <- 2
	c <- 3

	fmt.Println(<-c) //缓冲区尚有数据不会阻塞
	fmt.Println(<-c)
	fmt.Println(<-c)

	//缓冲区大小仅是内部属性，不属于类型组成部分。另外通道变量本身就是指针，可用相等操作符判断是否为同一对象或者为nil。
	//内置函数cap和len返回缓冲区大小和当前缓冲区数量。
	//而对于同步通道则都返回0，据此可判断通道是同步还是异步。
	fmt.Println(c == nil, cap(c), len(c))

	c2 := make(chan int) //无缓存通道，必须配对goroutine

	go func() {
		c2 <- 4
	}()

	fmt.Println(<-c2)

}
