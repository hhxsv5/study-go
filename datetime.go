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

	//时区转换，比如打开数据库连接时指定好时区：loc=Asia%2FShanghai
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-01-08 14:03:03", loc)
	fmt.Println(t.Unix(), t.String())
}
