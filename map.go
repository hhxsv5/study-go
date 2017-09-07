package main

import "fmt"

func main() {
	//Map 是一种无序的键值对的集合。通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。
	//Map 是无序的，无法决定它的返回顺序，因为 Map 是使用 hash 表来实现的。

	//申明string=>int的map1
	var map1 map[string]int
	//初始化map1
	map1 = map[string]int{"a": 1, "b": 2}
	fmt.Println(map1)

	//申明并初始化
	map2 := map[string]string{"a": "b", "c": "d"}
	fmt.Println(map2)

	//插入key-value
	map1["c"] = 3
	map1["d"] = 4

	for k, v := range map1 {
		fmt.Println("k=>", k, ", v=>", v)
	}

	//检查元素是否存在
	value, ok := map1["a"]
	fmt.Println("value=>", value, ", ok=>", ok)
	value, ok = map1["aa"]
	fmt.Println("value=>", value, ", ok=>", ok)

	//删除元素a和b
	delete(map1, "a")
	delete(map1, "b")
	fmt.Println(map1)

}
