package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Id    uint32 //大写字母开头的为public，外部可见
	Name  string
	Age   uint8
	price float32 //小写字母开头的为protected，整个package可见
}

func NewDog(id uint32, name string, age uint8, price float32) *Dog {
	return &Dog{id, name, age, price}
}

func main() {

	//Json Encode
	dog1 := NewDog(1, "狗狗", 10, 1000.00)
	json1, _ := json.Marshal(dog1)
	jsonStr1 := string(json1)
	fmt.Println(jsonStr1) //price不可见

	//Json Decode
	var dog2 Dog
	err := json.Unmarshal([]byte(jsonStr1), &dog2)
	fmt.Println(err, dog2)
}
