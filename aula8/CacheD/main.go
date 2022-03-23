package main

import "cached/cached"

func main() {
	server := cached.NewServer()
	server.Start()
}
