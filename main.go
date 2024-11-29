package main

import (
	"boilerplate_golangfront/handlers"
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
