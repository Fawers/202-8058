package main

import (
	"fmt"
	"time"
)

type Ticker struct {
	C       chan uint
	value   uint
	stopper chan bool
}

func NewTicker(initialValue uint) *Ticker {
	ticker := new(Ticker)
	ticker.C = make(chan uint)
	ticker.value = initialValue
	ticker.stopper = make(chan bool)

	go ticker.run()

	return ticker
}

func (t *Ticker) run() {
	running := true
	for running {
		select {
		case t.C <- t.value:
			t.value++

		case <-t.stopper:
			running = false
		}

		time.Sleep(time.Second)
	}
}

func (t *Ticker) Stop() {
	t.stopper <- true
}

func (t *Ticker) Reset(initialValue uint) {
	t.Stop()
	t.value = initialValue
	go t.run()
}

// timer:
// esperar DUR antes de enviar um sinal ao canal
// ou
// deixar de esperar se chamarmos Stop() antes
// nosso timer: chan bool
// true -> se esperou o tempo proposto
// false -> se foi interrompido antes

type Timer struct {
	C, stopper chan bool
	seconds    uint
}

func NewTimer(seconds uint) *Timer {
	timer := Timer{
		C:       make(chan bool),
		stopper: make(chan bool),
		seconds: seconds,
	}
	go timer.run()
	return &timer
}

func (t *Timer) run() {
	ticker := NewTicker(0)
	running := true
	for running {
		select {
		case v := <-ticker.C:
			if v >= t.seconds {
				t.C <- true
				running = false
			}

		case <-t.stopper:
			t.C <- false
			ticker.Stop()
			running = false
		}
	}
}

func (t *Timer) Stop() {
	t.stopper <- true
}

func (t *Timer) Reset(seconds uint) {
	t.Stop()
	t.seconds = seconds
	go t.run()
}

func main() {
	t := NewTimer(10)

	fmt.Println("esperando 10 segundos")

	go func() {
		time.Sleep(5 * time.Second)
		t.Stop()
	}()

	if <-t.C {
		fmt.Println("10 segundos aguardados com sucesso!!")
	} else {
		fmt.Println("algo interrompeu o timer")
	}
}

func main2() {
	t := NewTicker(0)

	for i := range t.C {
		if i >= 10 {
			t.Stop()
			break
		}

		fmt.Printf("recebi o valor %d do ticker\n", i)
	}
}
