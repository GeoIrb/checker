package app

import (
	"time"
)

type tick struct {
	Next chan struct{}
	Time time.Duration
}

//TickInit инициализирует работу тикера
func TickInit(d time.Duration) tick {
	tick := tick{
		Next: make(chan struct{}, 1),
		Time: d,
	}
	tick.Next <- struct{}{}
	return tick
}

//Wait ожидание
func (t tick) Wait() {
	go func() {
		time.Sleep(t.Time)
		t.Next <- struct{}{}
	}()
}
