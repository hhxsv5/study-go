package main

import "math/rand"
import "time"

func main() {
	//使用时间来作为随机数种子
	println("unix time", time.Now().Unix())
	println("unix nano time", time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	x := rnd.Intn(100) - rnd.Intn(100)

	if x > 0 {
		println("x", x)
	} else if x < 0 {
		println("-x", x)
	} else {
		println("0", x)
	}
}
