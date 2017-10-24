package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	currentTime := time.Now().Local()
	fmt.Println(currentTime)

	newFormat := currentTime.Format("2006-01-02 15:04:05.0000") //2006-01-02 15:04:05代替了PHP中的Y-m-d H:i:s
	fmt.Println(newFormat)

	//string to time
	timestr := "2011-06-06 00:00:00"
	mytime, _ := time.Parse("2006-01-02 15:04:05", timestr)
	fmt.Println(mytime, mytime.Unix())

	//计算执行耗时
	var total int64 = 0
	var t1 time.Time = time.Now()
	var i int64 = 0
	for ; i < 1000000110; i++ {
		total += i * 2
	}
	//Since()
	fmt.Println(total, time.Since(t1))
	fmt.Println(total, reflect.TypeOf(t1))
}
