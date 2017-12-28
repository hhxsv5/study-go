package main

import (
	"github.com/robfig/cron"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strconv"
)

func main() {
	ch := make(chan bool)
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Every second", time.Now().UnixNano()/1000/1000)
		ioutil.WriteFile("/Users/dave/Documents/www/go/src/test/cron.log", []byte("\ntest "+strconv.FormatInt(time.Now().UnixNano()/1000/1000, 10)), os.ModePerm)
	})
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	defer c.Stop()

	<-ch

	//for _, e := range c.Entries() {
	//	fmt.Println(e.Schedule)
	//}
	//..
	// Funcs are invoked in their own goroutine, asynchronously.
	//...
	// Funcs may also be added to a running Cron
	//c.AddFunc("@daily", func() { fmt.Println("Every day") })
	//..
	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())
	//..
	//c.Stop()  // Stop the scheduler (does not stop any jobs already running).
}
