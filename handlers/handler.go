package handlers

import (
	"fmt"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

// Ini untuk handle pathname (istilah lainnya "routing")
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func RenderRegister(c *gin.Context) {
	c.HTML(200, "registrasi.html", gin.H{})
}

func HandleRegistrasi(c *gin.Context) {
	var (
		acc Mahasiswa
	)

	c.Bind(&acc)

	// Generate nomor antrian
	now := time.Now()
	formattedTime := now.Format("20060102")

	config := GetServer(c).Config // ambil config dari server

	antrianPrefix := fmt.Sprintf("antrian:%v", formattedTime)
	antrianData, err := getAllKeysWithPrefix(antrianPrefix)
	handleError(err)

	if len(antrianData) >= config.SASC.MaxMahasiswaHari {
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

	c.Redirect(302, "/")
}

func RenderHomeWithAntrian(c *gin.Context) {
	config := GetServer(c).Config // ambil config dari server

	antrianDataGrouped := make(map[string]map[string][]struct {
		NoAntrian     string
		SudahDilayani bool
		Nim           string
		Nama          string
		Tanggal       string
		DurasiKonsul  string
	})

	now := time.Now()
	for i := 0; i < 3; i++ {
		currentDate := now.AddDate(0, 0, i)
		formattedDate := currentDate.Format("20060102")
		antrianPrefix := fmt.Sprintf("antrian:%v", formattedDate)
		antrianData, err := getAllKeysWithPrefix(antrianPrefix)
		handleError(err)

		if len(antrianData) > 0 {
			antrianDataGrouped[currentDate.Format("02 January 2006")] = make(map[string][]struct {
				NoAntrian     string
				SudahDilayani bool
				Nim           string
				Nama          string
				Tanggal       string
				DurasiKonsul  string
			})
		}

		antrianList := []Antrian{}
		j := 0
		for _, value := range antrianData {
			antrianMap, ok := value.(map[string]interface{})
			if !ok {
				fmt.Println("Error type assertion", value)
				continue
			}
			var antrian Antrian

			for k, v := range antrianMap {
				switch k {
				case "NoAntrian":
					str, ok := v.(string)
					if !ok {
						fmt.Println("Error type asserting string", k)
						continue
					}
					antrian.NoAntrian = str
				case "Nim":
					str, ok := v.(string)
					if !ok {
						fmt.Println("Error type asserting string", k)
						continue
					}
					antrian.Nim = str
				case "SudahDilayani":
					boolean, ok := v.(bool)
					if !ok {
						fmt.Println("error type asserting bool", k)
						continue
					}
					antrian.SudahDilayani = boolean
				default:
				}
			}
			antrianList = append(antrianList, antrian)
			j++
		}

		for _, antrian := range antrianList {
			mahasiswaData, err := getFromDB(fmt.Sprintf("mahasiswa:%v", antrian.Nim))
			handleError(err)

			mahasiswaMap, ok := mahasiswaData.(map[string]interface{})
			if !ok {
				fmt.Println("Error type assertion", antrian.Nim)

				continue
			}
			var mahasiswa Mahasiswa
			for k, v := range mahasiswaMap {
				switch k {
				case "Nim":
					str, ok := v.(string)
					if !ok {
						fmt.Println("error type asserting string", k)
						continue
					}
					mahasiswa.Nim = str
				case "Nama":
					str, ok := v.(string)
					if !ok {
						fmt.Println("error type asserting string", k)

						continue
					}
					mahasiswa.Nama = str
				case "Jurusan":
					str, ok := v.(string)
					if !ok {
						fmt.Println("Error type asserting string", k)

						continue
					}
					mahasiswa.Jurusan = str
				case "JenisKonsul":
					str, ok := v.(string)
					if !ok {
						fmt.Println("Error type asserting string", k)

						continue
					}
					mahasiswa.JenisKonsul = str
				default:
				}
			}
			startTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 9, 0, 0, 0, time.UTC)
			durasiMulai := startTime.Add(time.Duration(config.SASC.DurasiKonsultasi*j) * time.Minute)
			durasiSelesai := durasiMulai.Add(time.Duration(config.SASC.DurasiKonsultasi) * time.Minute)
			mahasiswa.DurasiKonsul = fmt.Sprintf("%v-%v", durasiMulai.Format("15:04"), durasiSelesai.Format("15:04"))
			antrianDataGrouped[currentDate.Format("02 January 2006")][mahasiswa.Jurusan] = append(antrianDataGrouped[currentDate.Format("02 January 2006")][mahasiswa.Jurusan], struct {
				NoAntrian     string
				SudahDilayani bool
				Nim           string
				Nama          string
				Tanggal       string
				DurasiKonsul  string
			}{
				NoAntrian:     antrian.NoAntrian,
				SudahDilayani: antrian.SudahDilayani,
				Nim:           antrian.Nim,
				Nama:          mahasiswa.Nama,
				Tanggal:       currentDate.Format("02 January 2006"),
				DurasiKonsul:  mahasiswa.DurasiKonsul,
			})
			j++

		}

		for _, jurusanData := range antrianDataGrouped {
			for _, antrianList := range jurusanData {

				sort.Slice(antrianList, func(i, j int) bool {
					return antrianList[i].NoAntrian < antrianList[j].NoAntrian
				})
			}
		}
	}

	c.HTML(200, "index.html", gin.H{
		"antrian_data": antrianDataGrouped,
	})
}

func ApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server time: %v", time.Now().UTC().Format(time.RFC850)),
	})
}

// Helper function to get server instance from context
func GetServer(c *gin.Context) *Server {

	server, exists := c.Keys["server"].(*Server)

	if !exists {
		panic("Server instance not found in context")
	}
	return server
}
