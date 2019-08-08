package main

import (
	"fmt"
	"geoirb/checker/app"
	"geoirb/checker/handling"
	"time"
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
