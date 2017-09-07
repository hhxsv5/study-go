package main

//import "fmt"
//import "math/rand"

import ( //打包导入
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("hello golang")
	fmt.Println("Random number:", rand.Intn(100))
	fmt.Println("pi", math.Pi) //大写开头是导出的，外部才可以使用，math.pi报错
}
