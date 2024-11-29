package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Ini untuk handle pathname (istilah lainnya "routing")
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func RenderLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func RenderRegister(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func RenderGuest(c *gin.Context) {
	c.HTML(200, "guest.html", gin.H{})
}

func RenderHistory(c *gin.Context) {
	c.HTML(200, "history.html", gin.H{})
}

func RenderProfile(c *gin.Context) {
	c.HTML(200, "profile.html", gin.H{})
}

func RenderSecurity(c *gin.Context) {
	c.HTML(200, "security.html", gin.H{})
}

func RenderSettings(c *gin.Context) {
	c.HTML(200, "settings.html", gin.H{})
}

func HandleRegistrasi(c *gin.Context) {
	var (
		acc Account
	)

	c.Bind(&acc)

	// tes print data dari form ke console
	templatePrint := fmt.Sprintf("Nama: %v\nEmail: %v\nUsername: %v\nPassword: %v\nNoHandphone: %v\n", acc.Nama, acc.Email, acc.Username, acc.Password, acc.NoHandphone)
	fmt.Println("--- Test Registrasi ---")
	fmt.Println(templatePrint)
	fmt.Println("--- --- ---")

	// tambah data ke DB
	writeToDB(fmt.Sprintf("acc.user.%v.pass", acc.Username), acc.Password)
	writeToDB(fmt.Sprintf("acc.user.%v.nama", acc.Username), acc.Nama)
	writeToDB(fmt.Sprintf("acc.user.%v.email", acc.Username), acc.Email)
	writeToDB(fmt.Sprintf("acc.user.%v.nohp", acc.Username), fmt.Sprintf("%v", acc.NoHandphone))

	c.JSON(200, gin.H{
		"Nama":        acc.Nama,
		"Email":       acc.Email,
		"Username":    acc.Username,
		"Password":    acc.Password,
		"NoHandphone": acc.NoHandphone,
	})
}

func ApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server time: %v", time.Now().UTC().Format(time.RFC850)),
	})
}
