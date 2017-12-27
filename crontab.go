package main

import (
	"github.com/robfig/cron"
	"fmt"
)

func main() {
	ch := make(chan bool)
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every second") })
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	<-ch
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
