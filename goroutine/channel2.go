package main

import "fmt"

func sum(a []int, c chan int) {
	s := 0
	for _, v := range a {
		s += v
	}

	c <- s //计算结果写入channel

}

func main() {
	a1 := []int{1, 3}
	a2 := []int{2, 4}

	c := make(chan int)

	go sum(a1, c)
	go sum(a2, c)

	x := <-c //从channel读取结果: 6
	y := <-c //4

	fmt.Println(x, y)
}
