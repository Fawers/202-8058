package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	id int
	wg *sync.WaitGroup
	m  *sync.Mutex
}

func NewWorker(id int, wg *sync.WaitGroup, m *sync.Mutex) *Worker {
	wg.Add(1)
	return &Worker{
		id: id,
		wg: wg,
		m:  m,
	}
}

func (w *Worker) work(counter *int) {
	fmt.Printf("Worker[%d] começou o seu trabalho\n", w.id)
	time.Sleep(time.Duration(rand.Intn(1000)) * 10 * time.Millisecond)
	w.m.Lock()
	*counter--
	fmt.Printf("Worker[%d]: Hora de ir pra casa! (%d)\n", w.id, *counter)
	w.m.Unlock()
	w.wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var counter int

	m.Lock()
	for i := 0; i < 10; i++ {
		counter++
		go NewWorker(i, &wg, &m).work(&counter)
	}
	m.Unlock()

	fmt.Println("antes de wait")
	wg.Wait()
}
