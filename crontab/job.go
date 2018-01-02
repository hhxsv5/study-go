package main

import (
	"github.com/robfig/cron"
	"fmt"
	"time"
)

type SecondJob struct {
}

func (job SecondJob) Run() {
	fmt.Println("SecondJob", time.Now().Format("2006-01-02 15:04:05"))
}

type MinuteJob struct {
}

func (job MinuteJob) Run() {
	fmt.Println("MinuteJob", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	c := cron.New()

	c.AddJob("* * * * * *", SecondJob{})
	c.AddJob("0 * * * * *", MinuteJob{})

	c.Start()

	defer c.Stop() //关闭着计划任务，但是不能关闭已经在执行中的任务

	select {}
}
