package main

import (
	"4discovery/daemon"
	"fmt"
	"os"
)

func main() {
	daemon, err := daemon.New(":8083")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = daemon.Start(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
