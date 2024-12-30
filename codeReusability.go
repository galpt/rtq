package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/afero"
)

var (
	osFS        = afero.NewOsFs()
	lockAntrian = false
)

// fungsi-fungsi lockAntrian
func kunciAntrian() {
	lockAntrian = true
}
func bukaAntrian() {
	lockAntrian = false
}

// string-to-byte (stb)
// fungsi untuk ubah string jadi byte
func stb(sebuahString string) []byte {
	return []byte(sebuahString)
}

// byte-to-string (bts)
// fungsi untuk ubah byte jadi string
func bts(sebuahByte []byte) string {
	return string(sebuahByte)
}

// print apapun
func printApapun(sebuahData any) {
	fmt.Println(fmt.Sprintf("\n%v\n", sebuahData))
}

// ini kalau terjadi error, maka program langsung exit
func handleErrorFatal(sebuahError error) {
	if sebuahError != nil {
		bukaAntrian()
		log.Println(fmt.Sprintf("Error fatal : %v", sebuahError))
		log.Fatal(sebuahError)
	}
}

// fungsi untuk strings.ReplaceAll() otomatis return string
func replaceSemua(sebuahString string, teksYangMauDiganti string, gantiDenganApa string) string {
	replace := strings.ReplaceAll(sebuahString, teksYangMauDiganti, gantiDenganApa)
	return replace
}

// fungsi untuk split string
func splitStr(sebuahString string, pemisahString string) []string {
	split := strings.Split(sebuahString, pemisahString)
	return split
}

// fungsi untuk split pengaturan.txt atau data_mahasiswa.txt
func splitFormatPengaturan(sebuahString string, namaData string) []string {
	split1 := splitStr(sebuahString, fmt.Sprintf("%v=[", namaData))
	// Validasi apakah split1 memiliki setidaknya 2 elemen
	if len(split1) < 2 {
		return []string{} // return slice kosong jika tidak ada elemen dengan index 1
	}

	split2 := splitStr(split1[1], "]")
	if len(split2) < 1 {
		return []string{} // return slice kosong jika tidak ada elemen dengan index 1
	}
	split3 := splitStr(split2[0], "|")
	return split3
}

// fungsi untuk cek apakah file sudah ada
func cekApakahFileAda(namaFile string) bool {
	cekNamaFile, err := afero.Exists(osFS, namaFile)
	handleErrorFatal(err)

	if cekNamaFile {
		return true
	} else {
		return false
	}
}

// fungsi untuk cek apakah folder sudah ada
func cekApakahFolderAda(namaFolder string) bool {
	cekFolder, err := afero.DirExists(osFS, namaFolder)
	handleErrorFatal(err)

	if cekFolder {
		return true
	} else {
		return false
	}
}

// fungsi untuk bikin folder & subfolder baru
func bikinFolderBaru(namaFolder string) bool {
	fmt.Println(fmt.Sprintf("\nMembuat folder %v\n", namaFolder))

	createFolderDB := osFS.MkdirAll(namaFolder, 0777)
	handleErrorFatal(createFolderDB)

	fmt.Println(fmt.Sprintf("\nBerhasil membuat folder %v\n", namaFolder))

	return true
}

// fungsi untuk bikin file baru
func bikinFileBaru(namaFile string) bool {

	// cek apakah file namaFile sudah ada atau ngga.
	// kalo misalnya gada, maka buatin baru
	if !cekApakahFileAda(namaFile) {

		fmt.Println(fmt.Sprintf("Membuat file %v", namaFile))

		bikinFile, err := osFS.Create(namaFile)
		if err != nil {
			log.Println(fmt.Sprintf("Error membuat file: %v", err))
			return false
		}

		closeFile := bikinFile.Close()
		handleErrorFatal(closeFile)

		fmt.Println(fmt.Sprintf("File %v berhasil dibuat", namaFile))

		return true

	} else {
		return false
	}

}

// fungsi untuk tulis data string ke suatu file
func tulisStringKeFile(namaFile string, dataString string) bool {

	// cek apakah file ada
	if cekApakahFileAda(namaFile) {

		tulisFile := afero.WriteFile(osFS, namaFile, stb(dataString), 0777)
		if tulisFile != nil {
			log.Println(fmt.Sprintf("Error menulis string ke file: %v", tulisFile))
			return false
		}

		// kalau ga error return 'true'
		return true
	} else {
		return false
	}

}

// fungsi untuk tulis data string ke suatu file
func tulisByteKeFile(namaFile string, dataBytes []byte) bool {

	// cek apakah file ada
	if cekApakahFileAda(namaFile) {

		tulisFile := afero.WriteFile(osFS, namaFile, dataBytes, 0777)
		if tulisFile != nil {
			log.Println(fmt.Sprintf("Error menulis byte ke file: %v", tulisFile))
			return false
		}

		// kalau ga error return 'true'
		return true
	} else {
		return false
	}

}

// fungsi untuk baca file dan return data dalam bentuk string
func bacaFileReturnString(namaFile string) string {

	// cek apakah file ada
	if cekApakahFileAda(namaFile) {

		bacaFile, err := afero.ReadFile(osFS, namaFile)
		handleErrorFatal(err)

		return string(bacaFile)

	} else {
		return ""
	}

}

// fungsi untuk baca file dan return data dalam bentuk string
func bacaFileReturnByte(namaFile string) []byte {

	// cek apakah file ada
	if cekApakahFileAda(namaFile) {

		bacaFile, err := afero.ReadFile(osFS, namaFile)
		handleErrorFatal(err)

		return bacaFile

	} else {
		return []byte("")
	}

}

// fungsi untuk baca folder dan return array of folder
func bacaFolder(namaFolder string) []os.FileInfo {
	cekFolder, err := afero.ReadDir(osFS, namaFolder)
	handleErrorFatal(err)
	return cekFolder
}

// fungsi untuk hapus file atau folder
func hapusSemua(sebuahLokasi string) bool {
	hapus := osFS.RemoveAll(sebuahLokasi)
	if hapus != nil {
		log.Println(fmt.Sprintf("Error menghapus file atau folder: %v", hapus))
		return false
	} else {
		return true
	}
}

// fungsi untuk dapatkan rincian hari tanggal
func getInfoTanggal(t time.Time) (string, int, time.Month, int) {
	hari := t.Weekday().String()
	tanggal := t.Day()
	bulan := t.Month()
	tahun := t.Year()
	return hari, tanggal, bulan, tahun
}

// fungsi untuk baca waktu reservasi
func bacaWaktuReservasi(namaFile string) (time.Time, error) {

	timeString := bacaFileReturnString(namaFile)
	if timeString == "" {
		return time.Time{}, fmt.Errorf("file waktu reservasi kosong")
	}
	parsedTime, err := time.Parse(timeFormat, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// fungsi cek apakah waktu reservasi sudah lebih dari 24 jam.
// jika ya berarti sudah ganti hari, maka bisa replace dengan data baru.
func bandingkanWaktuReservasi(waktuReservasi time.Time) bool {
	waktuSekarang := time.Now()
	waktuBerlalu := waktuSekarang.Sub(waktuReservasi)
	if waktuBerlalu > 24*time.Hour {
		return true
	} else {
		return false
	}
}

// fungsi untuk simpan waktu reservasi
func simpanWaktuReservasi(namaFile string) bool {

	// cek apakah file ada
	if !cekApakahFileAda(namaFile) {

		waktuSekarang := time.Now()
		timeString := waktuSekarang.Format(timeFormat)
		bikinFileBaru(namaFile)
		tulisStringKeFile(namaFile, timeString)

		// kalau ga error return 'true'
		return true
	} else {

		// jika file sudah ada maka cek apakah sudah lebih dari 24 jam
		getWaktu, err := bacaWaktuReservasi(namaFile)
		if err != nil {
			log.Println(fmt.Sprintf("Error membaca waktu reservasi: %v", err))
			return false
		}

		bolehReplace := bandingkanWaktuReservasi(getWaktu)
		if bolehReplace {

			hapusSemua(namaFile)

			waktuSekarang := time.Now()
			timeString := waktuSekarang.Format(timeFormat)
			bikinFileBaru(namaFile)
			tulisStringKeFile(namaFile, timeString)

			return true
		} else {
			return false
		}

	}

}
