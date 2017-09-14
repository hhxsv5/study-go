//当写自己包的时候，要使用短小的不含有 _(下划线)的小写单词来为文件命名。
package pack1

//public变量
var Pack1Int int32 = 10

//protected变量
var pack1Float float32 = 3.14

func Pack1Calc(x int32, y float32) float32 {
	return float32(x) * y
}
