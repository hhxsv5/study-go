//当写自己包的时候，要使用短小的不含有 _(下划线)的小写单词来为文件命名。
package pack

import "fmt"

//public变量
var PackInt int32 = 10

//protected变量
var packFloat float32 = 3.14

func PackCalc(x int32, y float32) float32 {
	return float32(x) * y
}

//每个源文件都可以定义一个或多个初始化函数。
//编译器不保证多个初始化函数执行次序。
//初始化函数在单一线程被调用，仅执行一次。
//初始化函数在包所有全局变量初始化后执行。
//在所有初始化函数结束后才执行 main。
//无法调用初始化函数。
//因为无法保证初始化函数执行顺序，因此全局变量应该直接用 var 初始化。
func init() {
	fmt.Println("run init1")
}

func init() {
	fmt.Println("run init2")
}

//可在初始化函数中使用 goroutine，可等待其结束。
func init() {
	c := make(chan bool)

	go func() {
		c <- true
	}()

	<-c
}
