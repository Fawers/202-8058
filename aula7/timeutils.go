package main

import (
	"fmt"
	"time"
)

func doTicker() *time.Ticker {
	ticker := time.NewTicker(time.Duration(time.Second))

	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	return ticker
}

func doTimer() {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println("timer de 5 segundos")
	t := <-timer.C
	fmt.Println(t)
	fmt.Println("acabou")
}

func main() {
	ticker := doTicker()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ticker.C)
	}

	ticker.Stop()
}
