package main

import "fmt"

func main() {
	//申明定长数组，其成员为unint32，默认初始化为0
	var userIds1 [5]uint32
	fmt.Println(userIds1)

	//申明并初始化定长数组
	var userIds2 = [5]uint32{1, 2, 3, 4, 5}
	fmt.Println(userIds2)

	//申明数组，并根据元素个数来确定长度
	var userIds3 = [...]uint32{1, 2, 3}
	//修改数组
	userIds3[2] = 4
	//访问数组元素
	fmt.Println(userIds3[2])

	//多维数组示例，二维
	var userGroups1 = [2][3]int32{
		{1, 2, 3},
		{4, 5, 6},
	}
	//访问二维数组
	fmt.Println(userGroups1[0][2])
	fmt.Println(userGroups1)
}
