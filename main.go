package main

import (
	"fmt"
	"time"

	"github.com/checker/app"
	"github.com/checker/handling"
)

func main() {
	step := app.TickInit(time.Duration(app.Load("time")["sleep"].(int)) * time.Second)
	for {
		select {
		case <-step.Next:
			conn := app.Init()

			startTime := time.Now()
			handling.Start(conn)

			fmt.Printf("\nTime to check %v\n\n", time.Now().Sub(startTime))
			conn.Pause()
		}
		step.Wait()
	}
}
