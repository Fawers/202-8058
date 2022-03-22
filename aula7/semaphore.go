package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	Name               string
	Status             string
	Red, Yellow, Green time.Duration
	timer              *time.Timer
	m                  *sync.Mutex
}

func NewSemaphore(name string, red, yellow, green time.Duration, m *sync.Mutex) *Semaphore {
	return &Semaphore{
		Name:   name,
		Status: "vermelho",
		Red:    red,
		Yellow: yellow,
		Green:  green,
		timer:  time.NewTimer(0),
		m:      m,
	}
}

func (s *Semaphore) Work() {
	s.timer.Stop()
	for {
		s.m.Lock()
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
		s.m.Unlock()
		<-s.timer.C
	}
}

func main() {
	var mutex sync.Mutex
	sems := []*Semaphore{
		NewSemaphore("Rua", 6*time.Second, time.Second, 3*time.Second, &mutex),
		NewSemaphore("Avenida", 3*time.Second, time.Second, 6*time.Second, &mutex),
	}

	for _, s := range sems {
		go s.Work()
	}

	time.Sleep(60 * time.Second)
}
