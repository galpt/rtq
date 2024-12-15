package main

import (
	"rtq/handlers"
)

const (
	listeningPort = ":8080"
)

var (
	server = handlers.Server{}
)

func main() {
	server.Init(listeningPort)
}
