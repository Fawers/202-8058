package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	Name               string
	Status             string
	Red, Yellow, Green time.Duration
	timer              *time.Timer
	previous, next     chan bool
}

func NewSemaphore(name string, red, yellow, green time.Duration, p, n chan bool) *Semaphore {
	s := Semaphore{
		Name:     name,
		Status:   "vermelho",
		Red:      red,
		Yellow:   yellow,
		Green:    green,
		timer:    nil,
		previous: p,
		next:     n,
	}

	return &s
}

func (s *Semaphore) Work() {
	s.timer = time.NewTimer(s.Red)

	for {
		<-s.previous

		s.Status = "verde"
		fmt.Printf("%s ficou %s\n", s.Name, s.Status)
		s.timer.Reset(s.Green)
		<-s.timer.C

		s.Status = "amarelo"
		fmt.Printf("%s ficou %s\n", s.Name, s.Status)
		s.timer.Reset(s.Yellow)
		<-s.timer.C

		s.Status = "vermelho"
		fmt.Printf("%s ficou %s\n", s.Name, s.Status)
		s.timer.Reset(s.Red)
		s.next <- true
		<-s.timer.C
	}
}

func main() {
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	sems := []*Semaphore{
		NewSemaphore("Rua", 6*time.Second, time.Second, 5*time.Second, a, b),
		NewSemaphore("Avenida", 3*time.Second, time.Second, 6*time.Second, b, c),
		NewSemaphore("Alameda", 3*time.Second, time.Second, 8*time.Second, c, a),
	}

	for _, s := range sems {
		go s.Work()
	}

	a <- true
	time.Sleep(60 * time.Second)
}
