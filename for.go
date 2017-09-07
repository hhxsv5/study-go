package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		if i == 20 {
			continue
		}
		if i == 21 {
			goto GOHERE
		}
		if i == 50 {
			fmt.Println("break 50")
			break
		}

		fmt.Println(i)
	}

	for true {
		fmt.Println("for loop", time.Now().Unix(), time.Second)
		time.Sleep(time.Second)
	}

GOHERE:
	fmt.Println("goto here")

}
