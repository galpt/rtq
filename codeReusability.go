package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/afero"
)

var (
	osFS = afero.NewOsFs()
)

// print apapun
func printApapun(sebuahData any) {
	fmt.Println(fmt.Sprintf("\n%v\n", sebuahData))
}

// ini kalau terjadi error, maka program langsung exit
func handleErrorFatal(sebuahError error) {
	if sebuahError != nil {
		log.Fatal(sebuahError)
	}
}

// fungsi untuk split string
func splitStr(sebuahString string, pemisahString string) []string {
	split := strings.Split(sebuahString, pemisahString)
	return split
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

	createFolderDB := osFS.MkdirAll(folderDB, 0777)
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

		bikinFile, err := osFS.Create(namaFilePengaturan)
		handleErrorFatal(err)

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
		bukaFile, err := osFS.Open(namaFile)
		handleErrorFatal(err)

		writeStatus, err := bukaFile.WriteString(dataString)
		handleErrorFatal(err)
		printApapun(writeStatus)

		tutupFile := bukaFile.Close()
		handleErrorFatal(tutupFile)

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
		bukaFile, err := osFS.Open(namaFile)
		handleErrorFatal(err)

		writeStatus, err := bukaFile.Write(dataBytes)
		handleErrorFatal(err)
		printApapun(writeStatus)

		tutupFile := bukaFile.Close()
		handleErrorFatal(tutupFile)

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
