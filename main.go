package main

import (
	"fmt"
	"rtq/handlers"

	"github.com/gin-gonic/gin"
)

const (
	listeningPort = ":8080"
)

var (
	server = handlers.Server{}
)

// ServerRecover middleware
func ServerRecover(server *handlers.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("server", server)
		c.Next()
	}
}

func main() {

	server.Router = gin.New()                 // Inisiasi Router di sini.
	server.Router.Use(ServerRecover(&server)) // Terapkan middleware sebelum inisiasi server
	server.Init(listeningPort)                // jalankan *init* di sini.

	fmt.Printf("Server starting on port %s\n", listeningPort)
	err := server.Router.Run(listeningPort)
	if err != nil {

		panic(fmt.Sprintf("Failed to start the application on port %s", listeningPort))
	}
}
