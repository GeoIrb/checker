package main

import (
	"fmt"

	"log"
	"os"
	"time"

	"github.com/GeoIrb/app"
	"github.com/GeoIrb/checker/handling"
)

const VERSION string = "1.0.0"
const KEYWORDS string = "Проверка наличия пользовательских слов"
const SYSTEM string = "Проверка наличия системных слов"
const HASH string = "Проверка хеш сумм"

func main() {

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		log.Println("Версия: ", VERSION)

		log.Print("Описание: ")
		switch os.Args[0] {
		case "hash":
			log.Println(HASH)
		case "system":
			log.Println(SYSTEM)
		case "keywords":
			log.Println(KEYWORDS)
		}
		return
	}

	step := app.NewTick(time.Duration(app.Load("time")["sleep"].(int)) * time.Second)
	for {
		select {
		case <-step.Step:
			conn := app.Init()

			startTime := time.Now()
			handling.Start(conn)

			fmt.Printf("\nTime to check %v\n\n", time.Now().Sub(startTime))
			conn.Disconnect()
		}
		step.Wait()
	}
}
