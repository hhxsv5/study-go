package main

import (
	"fmt"
	"runtime"
)

func main() {

	c1 := make(chan func())
	c2 := make(chan func())

	runtime.GOMAXPROCS(2)

	//wg := sync.WaitGroup{}
	//wg.Add(1)

	go func() {
		defer close(c1)
		for i := 0; i < 5; i++ {
			t := i //为什么需要弄个临时变量t？
			c1 <- func() {
				fmt.Println("anonymous func: c1", t)
			}
		}
	}()

	go func() {
		defer close(c2)
		for i := 0; i < 5; i++ {
			t := i
			c2 <- func() {
				fmt.Println("anonymous func: c2", t)
			}
		}
	}()

	func() {
		//defer wg.Done()
		var (
			f1  func()
			f2  func()
			ok1 = false
			ok2 = false
		)
		for {
			if f1, ok1 = <-c1; ok1 {
				f1()
			}
			if f2, ok2 = <-c2; ok2 {
				f2()
			}

			if !ok1 && !ok2 {
				break
			}
		}
	}()

	//wg.Wait()
}
