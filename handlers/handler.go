package handlers

import (
	"fmt"
	"strconv"
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

func ApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server time: %v", time.Now().UTC().Format(time.RFC850)),
	})
}

func HandleRegistrasi(c *gin.Context) {
	var (
		acc Mahasiswa
	)

	c.Bind(&acc)

	// Generate nomor antrian
	now := time.Now()
	formattedTime := now.Format("20060102")
	maxAntrianData, err := getFromDB("config:maxantrian")
	handleError(err)

	maxAntrianStr, ok := maxAntrianData.(string) // Type assertion to get the string
	if !ok {
		c.JSON(400, gin.H{"message": "maxantrian is not string"})
	}

	maxAntrian, err := strconv.Atoi(maxAntrianStr)
	handleError(err)

	antrianPrefix := fmt.Sprintf("antrian:%v", formattedTime)
	antrianData, err := getAllKeysWithPrefix(antrianPrefix)
	handleError(err)

	if len(antrianData) >= maxAntrian {
		c.JSON(400, gin.H{"message": "Maaf pendaftaran antrian sudah penuh untuk hari ini"})
		return
	}

	nomorAntrian := fmt.Sprintf("%03d", len(antrianData)+1) // Misal: 001, 002, dst.

	// Create data antrian
	newAntrian := Antrian{
		NoAntrian:     nomorAntrian,
		Nim:           acc.Nim,
		SudahDilayani: false,
		WaktuAntri:    time.Now().Format(time.RFC3339),
	}

	// Simpan data ke DB
	writeToDB(fmt.Sprintf("mahasiswa:%v", acc.Nim), acc)
	writeToDB(fmt.Sprintf("antrian:%v:%v", formattedTime, nomorAntrian), newAntrian)

	c.JSON(200, gin.H{
		"message":       "Pendaftaran Berhasil",
		"nomor_antrian": nomorAntrian,
		"durasi":        "10:00-10:45", // harusnya ini ditaruh di logic
	})
}
