package main

import "fmt"

//切片是对数组的抽象
//数组的长度不可改变，切片则是动态数组，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

func main() {
	//申明：切片不需要说明长度
	var slice1 []int

	if len(slice1) == 0 {
		fmt.Println("slice1是空切片", len(slice1), cap(slice1))
	}

	//使用make申明并初始化切片：创建长度为5，容量为10的切片，cap>=len
	var slice2 []int = make([]int, 5, 10)
	//申明：或简写：slice2 := make([]int, 5, 10)

	if len(slice2) {
		fmt.Println("slice2是空切片", len(slice2), cap(slice2))
	} else {
		fmt.Println(slice2)
	}

	//初始化：[]表示是切片类型，len=cap=3
	slice1 = []int{0, 1, 2, 3}

	//初始化：从数组arr中取一段来初始化，startIndex:endIndex-1，startIndex和endIndex均可不指定，startIndex默认为0，endIndex默认为length
	arr := [6]int{0, 1, 2, 3, 4, 5}
	slice2 = arr[0:4] //取出索引0-3
	fmt.Println(slice1, slice2)

	//len()计算长度，cap()计算容量
	fmt.Println(len(slice1), cap(slice2))

	//截取切片
	fmt.Println(slice2[:2])  //取出索引0-1
	fmt.Println(slice2[2:3]) //取出索引2
	fmt.Println(slice2[2:])  //取出索引2-length

	//追加6、7、8、9
	slice2 = append(slice2, 6, 7, 8, 9)
	fmt.Println(slice2, len(slice2), cap(slice2))

	//复制
	var slice3 []int = make([]int, len(slice2), cap(slice2))
	copy(slice3, slice2)
	fmt.Println(slice3)

}
