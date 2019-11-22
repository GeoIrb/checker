package app

import (
	"os"
	"os/signal"
	"time"
)

//Tick таймер
type Tick struct {
	Step   chan struct{}
	Cancel chan os.Signal
	Time   time.Duration
}

//NewTick инициализирует работу таймера
func NewTick(d time.Duration) Tick {
	tick := Tick{
		Step:   make(chan struct{}, 1),
		Cancel: make(chan os.Signal),
		Time:   d,
	}

	signal.Notify(tick.Cancel, os.Kill)
	tick.Step <- struct{}{}
	return tick
}

//Wait ожидание таймера
func (t Tick) Wait() {
	go func() {
		time.Sleep(t.Time)
		t.Step <- struct{}{}
	}()
}

//Stop останавливает таймер
func (t Tick) Stop() {
	t.Cancel <- os.Kill
}
