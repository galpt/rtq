// server.go
package main

import (
	"fmt"
	"net/http"
	"strings"

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

	// untuk munculkan halaman antrian
	r.GET("/antrian", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(buatDaftarAntrian("")))
	})

	// untuk munculkan halaman daftar
	r.GET("/daftar", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(munculkanHalamanPendaftaran()))
	})

	// untuk handle user yang ingin daftar konsul
	r.POST("/daftar", func(c *gin.Context) {
		// parsing data dari form
		var form FormDaftarKonsul
		c.Bind(&form)

		action := c.PostForm("action")
		if action == "" {
			c.Data(http.StatusOK, "text/html; charset=utf-8", stb(simpanAntrian(form.Nim, strings.ToUpper(form.Nama), form.Jurusan, form.JamKonsul, strings.ToUpper(form.JenisKonsul), "")))
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", stb(simpanAntrian(form.Nim, strings.ToUpper(form.Nama), form.Jurusan, form.JamKonsul, strings.ToUpper(form.JenisKonsul), action)))
		}
	})

	// untuk munculkan halaman login admin
	r.GET("/admin/login", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(loginAdmin))
	})

	// untuk handle login admin
	r.POST("/admin/login", func(c *gin.Context) {
		var form struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}
		c.Bind(&form)

		// baca file users.txt
		usersData := bacaFileReturnString("users.txt")
		if usersData == "" {
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		// split data users.txt
		splitUsers := splitStr(usersData, "\n")
		for _, user := range splitUsers {
			splitUserData := splitStr(user, "|")
			if len(splitUserData) == 3 {
				username := splitUserData[0]
				password := splitUserData[1]
				// role := splitUserData[2]

				if form.Username == username && form.Password == password {
					c.Redirect(http.StatusFound, "/admin/dashboard")
					return
				}
			}
		}

		c.String(http.StatusUnauthorized, "Invalid credentials")
	})

	// untuk munculkan halaman dashboard admin
	r.GET("/admin/dashboard", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", stb(dashboardAdmin))
	})

	fmt.Println()
	fmt.Println(fmt.Sprintf("Server jalan di %v", listenAddr))
	r.Run(listenAddr)
}
