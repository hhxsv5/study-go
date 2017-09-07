package main

import "fmt"

func main() {
	// range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	// 在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range slice {
		fmt.Println("k=>", k, ", v=>", v)
	}

	for _, v := range slice {
		//_丢弃key
		fmt.Println("v=>", v)
	}

	//用于map
	kvs := map[string]string{"a": "1", "b": "2", "1": "a", "2": "b"}
	for k, v := range kvs {
		fmt.Println("k=>", k, ", v=>", v)
	}

	//用于枚举Unicode字符串：k是字符索引位置，v是字符的Unicode值
	for k, v := range "golang" {
		fmt.Println("k=>", k, ", v=>", v)
	}
}
