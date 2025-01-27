// daftarKonsul.go
package main

import (
	"fmt"
	"log"
	"regexp"
)

type FormDaftarKonsul struct {
	Nim         string `form:"nim"`
	Nama        string `form:"name"`
	Jurusan     string `form:"major"`
	JamKonsul   string `form:"jam_konsultasi"`
	JenisKonsul string `form:"jenis_konsultasi"`
}

// fungsi untuk munculkan halaman daftar konsul dan otomatis
// tampilkan jurusan dan waktu konsul yang diatur di 'pengaturan.exe'
func munculkanHalamanPendaftaran() string {

	var (
		semuaJurusanStr   = `<option value="" disabled>Pilih jurusan</option>`
		semuaJamKonsulStr = `<option value="" disabled>Pilih jam konsultasi</option>`
	)

	bacaPengaturanTxtString := bacaFileReturnString(namaFilePengaturan)
	splitJurusan := splitFormatPengaturan(bacaPengaturanTxtString, "jurusan")
	splitJamKonsul := splitFormatPengaturan(bacaPengaturanTxtString, "jamkonsul")

	if len(splitJurusan) == 0 {
		fmt.Println("Error: Tidak ada data jurusan di konfigurasi.")
		// Atau return string error ke pengguna
		return replaceSemua(daftar, "(REPLACE-INI-DENGAN-FORM-INPUT)", "<p>Konfigurasi tidak valid.</p>")
	}

	if len(splitJamKonsul) == 0 {
		fmt.Println("Error: Tidak ada data jam konsultasi di konfigurasi.")
		// Atau return string error ke pengguna
		return replaceSemua(daftar, "(REPLACE-INI-DENGAN-FORM-INPUT)", "<p>Konfigurasi tidak valid.</p>")

	}

	for idxJurusan := range splitJurusan {
		semuaJurusanStr += fmt.Sprintf(`<option value="%v">%v</option>`, splitJurusan[idxJurusan], splitJurusan[idxJurusan])
	}

	// untuk pilihan jam konsul harus muncul untuk hari ini, besok, lusa
	for i := 0; i < totalHari; i++ {

		if i == 0 {
			semuaJamKonsulStr += `<option value="" disabled>Hari Ini</option>`
			for idxJamKonsul := range splitJamKonsul {
				semuaJamKonsulStr += fmt.Sprintf(`<option value="hariini|%v">Hari Ini, %v</option>`, splitJamKonsul[idxJamKonsul], splitJamKonsul[idxJamKonsul])
			}
		} else if i == 1 {
			semuaJamKonsulStr += `<option value="" disabled>Besok</option>`
			for idxJamKonsul := range splitJamKonsul {
				semuaJamKonsulStr += fmt.Sprintf(`<option value="besok|%v">Besok, %v</option>`, splitJamKonsul[idxJamKonsul], splitJamKonsul[idxJamKonsul])
			}
		} else if i == 2 {
			semuaJamKonsulStr += `<option value="" disabled>Lusa</option>`
			for idxJamKonsul := range splitJamKonsul {
				semuaJamKonsulStr += fmt.Sprintf(`<option value="lusa|%v">Lusa, %v</option>`, splitJamKonsul[idxJamKonsul], splitJamKonsul[idxJamKonsul])
			}
		}
	}

	halamanPendaftaran := fmt.Sprintf(`
            <!-- NIM Input -->
            <div class="space-y-2">
                <label for="nim" class="block text-sm font-medium text-gray-700">
                    NIM (Nomor Induk Mahasiswa)
                </label>
                <input type="text" id="nim" name="nim"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan NIM" required>
            </div>

            <!-- Name Input -->
            <div class="space-y-2">
                <label for="name" class="block text-sm font-medium text-gray-700">
                    Nama
                </label>
                <input type="text" id="name" name="name"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan nama lengkap" required>
            </div>

            <!-- Major Select -->
            <div class="space-y-2">
                <label for="major" class="block text-sm font-medium text-gray-700">
                    Jurusan
                </label>
                <select id="major" name="major"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    required>
                    %v
                </select>
            </div>

            <!-- Pilih jam konsul -->
            <div class="space-y-2">
                <label for="jam_konsultasi" class="block text-sm font-medium text-gray-700">
                    Jam Konsultasi
                </label>
                <select id="jam_konsultasi" name="jam_konsultasi"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    required>
                    %v
                </select>
            </div>
			
			<!-- Pilih jenis konsultasi -->
            <div class="space-y-2">
                <label for="jenis_konsultasi" class="block text-sm font-medium text-gray-700">
                    Jenis Konsultasi
                </label>
                <select id="jenis_konsultasi" name="jenis_konsultasi"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    required>
                    <option value="" disabled>Pilih jenis konsultasi</option>
                    <option value="akademik">Akademik</option>
                    <option value="non-akademik">Non-akademik</option>
                </select>
            </div>
			`, semuaJurusanStr, semuaJamKonsulStr)

	return replaceSemua(daftar, "(REPLACE-INI-DENGAN-FORM-INPUT)", halamanPendaftaran)
}

// fungsi untuk cek apakah sudah ada antrian atau belum
func simpanAntrian(nim string, nama string, jurusan string, jamkonsulInput string, jeniskonsul string, action string) string {
	if action != "konfirmasi" {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	if lockAntrian {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	kunciAntrian()
	defer bukaAntrian() // Pastikan lock selalu dilepaskan

	if nim == "" {
		return buatDaftarAntrian("<div id=\"alert-2\" class=\"flex items-center p-4 mb-4 text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400\" role=\"alert\"><svg class=\"flex-shrink-0 w-4 h-4\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z\"/></svg><span class=\"sr-only\">Info</span><div class=\"ms-3 text-sm font-medium\">NIM wajib diisi. <a href=\"/daftar\" class=\"font-semibold underline hover:no-underline\">Kembali ke halaman daftar konsultasi</a>.</div></div>")
	}

	if nama == "" {
		return buatDaftarAntrian("<div id=\"alert-2\" class=\"flex items-center p-4 mb-4 text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400\" role=\"alert\"><svg class=\"flex-shrink-0 w-4 h-4\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z\"/></svg><span class=\"sr-only\">Info</span><div class=\"ms-3 text-sm font-medium\">Nama wajib diisi. <a href=\"/daftar\" class=\"font-semibold underline hover:no-underline\">Kembali ke halaman daftar konsultasi</a>.</div></div>")
	}

	match, _ := regexp.MatchString(`^\d+$`, nim)
	if !match {
		return buatDaftarAntrian("<div id=\"alert-2\" class=\"flex items-center p-4 mb-4 text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400\" role=\"alert\"><svg class=\"flex-shrink-0 w-4 h-4\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z\"/></svg><span class=\"sr-only\">Info</span><div class=\"ms-3 text-sm font-medium\">NIM harus berupa angka. <a href=\"/daftar\" class=\"font-semibold underline hover:no-underline\">Kembali ke halaman daftar konsultasi</a>.</div></div>")
	}

	folderHari := splitStr(jamkonsulInput, "|")
	lokasiFolder := ""
	jamkonsul := ""
	if len(folderHari) < 2 {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}
	jamkonsul = folderHari[1]
	if folderHari[0] == "hariini" {
		lokasiFolder = folderDBHariIni
	} else if folderHari[0] == "besok" {
		lokasiFolder = folderDBBesok
	} else if folderHari[0] == "lusa" {
		lokasiFolder = folderDBLusa
	}

	lokasiJurusan := lokasiFolder + fmt.Sprintf("/%v", jurusan)
	if !cekApakahFolderAda(lokasiJurusan) {
		log.Printf("Folder jurusan tidak ditemukan: %v\n", lokasiJurusan)
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	lokasiJamKonsul := lokasiJurusan + fmt.Sprintf("/%v", jamkonsul)
	if !cekApakahFolderAda(lokasiJamKonsul) {
		log.Printf("Folder jam konsul tidak ditemukan: %v\n", lokasiJamKonsul)
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	lokasiFileWaktuRsv := lokasiJamKonsul + fmt.Sprintf("/%v", namaFileWaktuReservasi)
	lokasiFileDataMhs := lokasiJamKonsul + fmt.Sprintf("/%v", namaFileDataMahasiswa)
	lokasiFileJenisKonsul := lokasiJamKonsul + fmt.Sprintf("/%v", namaFileJenisKonsul)

	// validasi apakah jurusan itu di jam segitu sudah ada yang reservasi atau tidak.
	// jika bisa reservasi maka simpan data terbaru ke lokasi folder tersebut.
	// jika tidak maka return pesan error.
	if cekApakahFileAda(lokasiFileWaktuRsv) {
		// jika ada file reservasi di lokasi tersebut maka cek
		// apakah sesi konsul sudah kadaluarsa atau belum
		getWaktu, err := bacaWaktuReservasi(lokasiFileWaktuRsv)
		if err != nil {
			log.Printf("Error membaca waktu reservasi: %v\n", err)
			return buatDaftarAntrian(alertGagalDaftarKonsul)
		}
		if !bandingkanWaktuReservasi(getWaktu) {
			// Jika belum kadaluarsa, tidak bisa daftar
			return buatDaftarAntrian(alertGagalDaftarKonsul)
		}
	}

	// Jika sudah kadaluarsa atau belum ada reservasi, lakukan pembaruan
	// hapus data lama
	hapusSemua(lokasiFileWaktuRsv)
	hapusSemua(lokasiFileDataMhs)
	hapusSemua(lokasiFileJenisKonsul)

	// tulis waktu reservasi baru
	simpanWaktuRsv := simpanWaktuReservasi(lokasiFileWaktuRsv)
	if !simpanWaktuRsv {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	// tulis data mahasiswa baru
	newFileDataMhs := bikinFileBaru(lokasiFileDataMhs)
	if !newFileDataMhs {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}
	tulisDataMhs := tulisStringKeFile(lokasiFileDataMhs, fmt.Sprintf("nimnama=[%v|%v]", nim, nama))
	if !tulisDataMhs {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	// tulis jenis konsul mahasiswa baru
	newFileJenisKonsul := bikinFileBaru(lokasiFileJenisKonsul)
	if !newFileJenisKonsul {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}
	tulisJenisKonsul := tulisStringKeFile(lokasiFileJenisKonsul, jeniskonsul)
	if !tulisJenisKonsul {
		return buatDaftarAntrian(alertGagalDaftarKonsul)
	}

	return buatDaftarAntrian(alertSuksesDaftarKonsul)
}
