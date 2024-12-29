package main

import (
	"fmt"
)

func cekPengaturan() {

	// kalo misalnya gada file "pengaturan.txt"
	// maka buatin file ini dengan isi dari template yang sudah kita tetapkan
	if !cekApakahFileAda(namaFilePengaturan) {

		fmt.Println("Membuat file pengaturan baru...")

		// step 1. bikin file baru
		bikinFileBaru(namaFilePengaturan)

		// step 2. isi file pengaturan dengan template default yang sudah ditentukan
		tulisStringKeFile(namaFilePengaturan, templateDefaultPengaturan)

		fmt.Println("File pengaturan berhasil dibuat")

	}

}

func bacaPengaturan() {

	var arrayHari = []string{folderDBHariIni, folderDBBesok, folderDBLusa}

	// kita baca dulu file "pengaturan.txt"
	bacaPengaturanTxtString := bacaFileReturnString(namaFilePengaturan)

	for idxHari := range arrayHari {

		// step 1. kalo misalnya folder 'db' tidak ada maka bikin baru
		if !cekApakahFolderAda(arrayHari[idxHari]) {

			// step 2. bikin folder jurusan di dalam folder db untuk hari ini, besok, lusa
			splitJurusan1 := splitStr(bacaPengaturanTxtString, "jurusan=[")
			splitJurusan2 := splitStr(splitJurusan1[1], "]")
			splitJurusan3 := splitStr(splitJurusan2[0], "|")

			splitJamKonsul1 := splitStr(bacaPengaturanTxtString, "jamkonsul=[")
			splitJamKonsul2 := splitStr(splitJamKonsul1[1], "]")
			splitJamKonsul3 := splitStr(splitJamKonsul2[0], "|")

			// looping sesuai jumlah jurusan yang dimasukkin
			for idxJurusan := range splitJurusan3 {

				bikinFolderBaru(arrayHari[idxHari] + fmt.Sprintf("/%v", splitJurusan3[idxJurusan]))

				// step 3. bikin folder jam konsul di dalam folder tiap jurusan
				// looping sesuai jumlah jurusan yang dimasukkin
				for idxJamKonsul := range splitJamKonsul3 {
					bikinFolderBaru(arrayHari[idxHari] + fmt.Sprintf("/%v/%v", splitJurusan3[idxJurusan], splitJamKonsul3[idxJamKonsul]))
				}
			}

		}
	}

}
