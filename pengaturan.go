package main

import (
	"fmt"
)

const (
	namaFilePengaturan    = "./pengaturan.txt"
	namaFileDataMahasiswa = "data_mahasiswa.txt"
	folderDB              = "./db"
)

var (
	templateDefaultPengaturan = `// harus dipisah dengan tanda '|'
jurusan=[Computer Science|Information Systems|International Relations]

// nama konselor harus berurutan dengan posisi 'jurusan' yang diisi di atas
namakonselor=[Ita|Ika|Agus]

// jam konsul sesuaikan dengan format seperti ini
jamkonsul=[09.00-09.45|09.45-10.30|10.30-11.15|11.15-12.00|13.00-13.45|13.45-14.30|14.30-15.15|15.15-16.00]`
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
	// kita baca dulu file "pengaturan.txt"
	bacaPengaturanTxtString := bacaFileReturnString(namaFilePengaturan)

	// step 1. kalo misalnya folder 'db' tidak ada maka bikin baru
	if !cekApakahFolderAda(folderDB) {
		bikinFolderBaru(folderDB)

		// step 2. bikin folder jurusan di dalam folder db
		splitJurusan1 := splitStr(bacaPengaturanTxtString, "jurusan=[")
		splitJurusan2 := splitStr(splitJurusan1[1], "]")
		splitJurusan3 := splitStr(splitJurusan2[0], "|")

		splitJamKonsul1 := splitStr(bacaPengaturanTxtString, "jamkonsul=[")
		splitJamKonsul2 := splitStr(splitJamKonsul1[1], "]")
		splitJamKonsul3 := splitStr(splitJamKonsul2[0], "|")

		// looping sesuai jumlah jurusan yang dimasukkin
		for idxJurusan := range splitJurusan3 {

			bikinFolderBaru(folderDB + fmt.Sprintf("/%v", splitJurusan3[idxJurusan]))

			// step 3. bikin folder jam konsul di dalam folder tiap jurusan
			// looping sesuai jumlah jurusan yang dimasukkin
			for idxJamKonsul := range splitJamKonsul3 {
				bikinFolderBaru(folderDB + fmt.Sprintf("/%v/%v", splitJurusan3[idxJurusan], splitJamKonsul3[idxJamKonsul]))
			}
		}

	}

}
