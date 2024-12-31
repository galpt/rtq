package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// fungsi untuk auto-format table per-jurusan
func buatTablePerJurusan(lokasiFolder string, namaJurusan string, hariTanggal string) string {
	// template untuk table header
	templateNamaJurusan := fmt.Sprintf(`<div
                class="text-center font-bold text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 py-3">
                %v
            </div>`, namaJurusan)

	// dapatkan nama konselor
	bacaPengaturanTxtString := bacaFileReturnString(namaFilePengaturan)
	splitJurusan := splitFormatPengaturan(bacaPengaturanTxtString, "jurusan")
	splitKonselor := splitFormatPengaturan(bacaPengaturanTxtString, "namakonselor")
	namaKonselor := ""
	for idxKonselor := range splitJurusan {
		if splitJurusan[idxKonselor] == namaJurusan {
			namaKonselor = splitKonselor[idxKonselor]
			break
		}
	}

	// template untuk per-baris
	var arrayTemplatePerBaris []string

	isiFolderJurusan := bacaFolder(lokasiFolder + fmt.Sprintf("/%v", namaJurusan))
	if len(isiFolderJurusan) == 0 {
		return fmt.Sprintf("<p>Belum ada antrian pada jurusan %v</p>", namaJurusan)
	}

	for idxJamKonsul := range isiFolderJurusan {
		jamKonsulDir := isiFolderJurusan[idxJamKonsul]
		if !jamKonsulDir.IsDir() {
			continue // Skip jika bukan directory
		}

		jamKonsulName := jamKonsulDir.Name()
		lokasiJamKonsul := lokasiFolder + fmt.Sprintf("/%v/%v", namaJurusan, jamKonsulName)

		// cek apakah folder jam konsul ada
		if !cekApakahFolderAda(lokasiJamKonsul) {
			log.Printf("Folder jam konsul tidak ditemukan: %v\n", lokasiJamKonsul)
			continue
		}

		// baca file 'data_mahasiswa.txt' dari tiap folder jam konsul
		var (
			mhsNim               = "Belum ada antrian"
			mhsNama              = "Belum ada antrian"
			mhsJenisKonsul       = "Belum ada antrian"
			lokasiDataMhs        = lokasiJamKonsul + fmt.Sprintf("/%v", namaFileDataMahasiswa)
			lokasiWaktuRsvMhs    = lokasiJamKonsul + fmt.Sprintf("/%v", namaFileWaktuReservasi)
			lokasiJenisKonsulMhs = lokasiJamKonsul + fmt.Sprintf("/%v", namaFileJenisKonsul)
		)

		if !cekApakahFileAda(lokasiWaktuRsvMhs) {
			// File waktu reservasi tidak ada, berarti belum ada antrian
			// Nilai default sudah ditetapkan
		} else {
			// baca file "waktu_reservasi.txt"
			waktuReservasi, err := bacaWaktuReservasi(lokasiWaktuRsvMhs)
			if err != nil {
				log.Printf("Error membaca waktu reservasi: %v\n", err)
				// Nilai default sudah ditetapkan
			} else {
				if bandingkanWaktuReservasi(waktuReservasi) {
					// jika waktu reservasi mahasiswa sudah lebih dari 24 jam,
					// maka sudah kadaluarsa dan hapus semua file terkait
					hapusSemua(lokasiWaktuRsvMhs)
					hapusSemua(lokasiDataMhs)
					hapusSemua(lokasiJenisKonsulMhs)
					// Nilai default untuk mhsNim, mhsNama, dan mhsJenisKonsul tetap "Belum ada antrian"
				} else {
					// baca file "data_mahasiswa.txt"
					dataMhs := bacaFileReturnString(lokasiDataMhs)

					if len(dataMhs) > 0 {
						splitDataMhs := splitFormatPengaturan(dataMhs, "nimnama")
						if len(splitDataMhs) >= 2 {
							mhsNim = splitDataMhs[0]
							mhsNama = splitDataMhs[1]

							// baca file "jenis_konsul.txt"
							mhsJenisKonsul = bacaFileReturnString(lokasiJenisKonsulMhs)
							if len(mhsJenisKonsul) < 1 {
								// Nilai default sudah ditetapkan
							}
						}
					}
				}
			}
		}

		templatePerBaris := fmt.Sprintf(`<tr
                            class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                            <th scope="row"
                                class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                %v
                            </th>
                            <td class="px-6 py-4">
                                %v
                            </td>
                            <td class="px-6 py-4">
                                %v
                            </td>
                            <td class="px-6 py-4">
                                %v
                            </td>
                            <td class="px-6 py-4">
                                %v
                            </td>
							<td class="px-6 py-4">
                                %v
                            </td>
                        </tr>`, hariTanggal, jamKonsulName, namaKonselor, mhsNim, mhsNama, mhsJenisKonsul)

		arrayTemplatePerBaris = append(arrayTemplatePerBaris, templatePerBaris)

	}

	templateSeluruhKolom := ""
	for idxBaris := range arrayTemplatePerBaris {
		templateSeluruhKolom += arrayTemplatePerBaris[idxBaris]
	}

	templateTablePerJurusan := fmt.Sprintf(`<div class="flex flex-col rounded-lg overflow-hidden">
            %v
			
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">
                                Hari Tanggal
                            </th>
                            <th scope="col" class="px-6 py-3">
                                Jam Konsultasi
                            </th>
                            <th scope="col" class="px-6 py-3">
                                Konselor
                            </th>
                            <th scope="col" class="px-6 py-3">
                                NIM
                            </th>
                            <th scope="col" class="px-6 py-3">
                                Nama
                            </th>
							<th scope="col" class="px-6 py-3">
                                Jenis Konsultasi
                            </th>
                        </tr>
                    </thead>
					
                    <tbody>
                        %v
                    </tbody>
                </table>
            </div>
        </div>`, templateNamaJurusan, templateSeluruhKolom)

	return templateTablePerJurusan
}

// fungsi untuk auto-generate table daftar antrian
func buatDaftarAntrian(notifikasi string) string {

	timenow := time.Now()

	hariIni, hariIniTgl, hariIniBln, hariIniThn := getInfoTanggal(timenow)

	besok := timenow.AddDate(0, 0, 1)
	hariBesok, hariBesokTgl, hariBesokBln, hariBesokThn := getInfoTanggal(besok)

	lusa := timenow.AddDate(0, 0, 2)
	hariLusa, hariLusaTgl, hariLusaBln, hariLusaThn := getInfoTanggal(lusa)

	templateSeluruhTable := notifikasi

	var arrayHari = []string{folderDBHariIni, folderDBBesok, folderDBLusa}

	// susun table daftar antrian
	for i := 0; i < totalHari; i++ {
		var daftarJurusan []os.FileInfo
		if i == 0 {
			daftarJurusan = bacaFolder(arrayHari[i])
		} else if i == 1 {
			daftarJurusan = bacaFolder(arrayHari[i])
		} else if i == 2 {
			daftarJurusan = bacaFolder(arrayHari[i])
		}
		if len(daftarJurusan) == 0 {
			if i == 0 {
				templateSeluruhTable += `<h4
                    class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
                    Antrian <span
                        class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">hari
                        ini</span></h4>`
			} else if i == 1 {
				templateSeluruhTable += `<h4
                    class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
                    Antrian <span
                        class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">besok</span></h4>`
			} else if i == 2 {
				templateSeluruhTable += `<h4
                    class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
                    Antrian <span
                        class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">lusa</span></h4>`
			}
			templateSeluruhTable += fmt.Sprintf("<p>Belum ada data antrian untuk %v</p>", arrayHari[i])
			continue
		}

		if i == 0 {
			templateSeluruhTable += `<h4
            class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
            Antrian <span
                class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">hari
                ini</span></h4>`
		} else if i == 1 {
			templateSeluruhTable += `<h4
            class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
            Antrian <span
                class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">besok</span></h4>`
		} else if i == 2 {
			templateSeluruhTable += `<h4
            class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-black">
            Antrian <span
                class="underline underline-offset-3 decoration-8 decoration-blue-400 dark:decoration-blue-600">lusa</span></h4>`
		}

		for idxJurusan := range daftarJurusan {
			if i == 0 {
				templateSeluruhTable += buatTablePerJurusan(folderDBHariIni, daftarJurusan[idxJurusan].Name(), fmt.Sprintf("%v, %v %v %v", hariIni, hariIniTgl, hariIniBln, hariIniThn))
			} else if i == 1 {
				templateSeluruhTable += buatTablePerJurusan(folderDBBesok, daftarJurusan[idxJurusan].Name(), fmt.Sprintf("%v, %v %v %v", hariBesok, hariBesokTgl, hariBesokBln, hariBesokThn))
			} else if i == 2 {
				templateSeluruhTable += buatTablePerJurusan(folderDBLusa, daftarJurusan[idxJurusan].Name(), fmt.Sprintf("%v, %v %v %v", hariLusa, hariLusaTgl, hariLusaBln, hariLusaThn))
			}

			templateSeluruhTable += "<br>"
		}
	}

	return replaceSemua(antrian, "(REPLACE-INI-DENGAN-TABLE)", templateSeluruhTable)

}
