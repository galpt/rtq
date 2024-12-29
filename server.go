package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	listenAddr = "0.0.0.0:8080"
)

func startServer() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())

	// untuk munculkan halaman home
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(home))
	})

	// untuk munculkan halaman daftar
	r.GET("/daftar", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(daftar))
	})

	// untuk munculkan halaman antrian
	r.GET("/antrian", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(buatDaftarAntrian()))
	})

	fmt.Println()
	fmt.Println(fmt.Sprintf("Server jalan di %v", listenAddr))
	r.Run(listenAddr)
}