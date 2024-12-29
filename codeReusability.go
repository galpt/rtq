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
	osFS = afero.NewOsFs()
)

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
		log.Fatal(sebuahError)
	}
}

// fungsi untuk split string
func splitStr(sebuahString string, pemisahString string) []string {
	split := strings.Split(sebuahString, pemisahString)
	return split
}

// fungsi untuk split pengaturan.txt atau data_mahasiswa.txt
func splitFormatPengaturan(sebuahString string, namaData string) []string {
	split1 := splitStr(sebuahString, fmt.Sprintf("%v=[", namaData))
	split2 := splitStr(split1[1], "]")
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

		tulisFile := afero.WriteFile(osFS, namaFile, stb(dataString), 0777)
		handleErrorFatal(tulisFile)

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
		handleErrorFatal(tulisFile)

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

// fungsi untuk dapatkan rincian hari tanggal
func getInfoTanggal(t time.Time) (string, int, time.Month, int) {
	hari := t.Weekday().String()
	tanggal := t.Day()
	bulan := t.Month()
	tahun := t.Year()
	return hari, tanggal, bulan, tahun
}

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
	for idxJurusan := range isiFolderJurusan {

		// baca file 'data_mahasiswa.txt' dari tiap folder jam konsul
		var (
			mhsNim  = ""
			mhsNama = ""
		)

		dataMhs := bacaFileReturnString(lokasiFolder + fmt.Sprintf("/%v/%v", isiFolderJurusan[idxJurusan].Name(), namaFileDataMahasiswa))

		if len(dataMhs) < 1 {
			mhsNim = "Belum ada antrian"
			mhsNama = "Belum ada antrian"
		} else {
			splitDataMhs := splitFormatPengaturan(dataMhs, "nimnama")
			mhsNim = splitDataMhs[0]
			mhsNama = splitDataMhs[1]
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
                        </tr>`, hariTanggal, isiFolderJurusan[idxJurusan].Name(), namaKonselor, mhsNim, mhsNama)

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
                                NAMA
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
func buatDaftarAntrian() string {

	timenow := time.Now()

	hariIni, hariIniTgl, hariIniBln, hariIniThn := getInfoTanggal(timenow)

	besok := timenow.AddDate(0, 0, 1)
	hariBesok, hariBesokTgl, hariBesokBln, hariBesokThn := getInfoTanggal(besok)

	lusa := timenow.AddDate(0, 0, 2)
	hariLusa, hariLusaTgl, hariLusaBln, hariLusaThn := getInfoTanggal(lusa)

	templateSeluruhTable := ""
	templateSeluruhHalaman := ""

	// ngecek di dalem folder 'db' ada jurusan apa aja har ini
	daftarJurusan := bacaFolder(folderDBHariIni)

	// susun table daftar antrian
	for i := 0; i < totalHari; i++ {

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
				templateSeluruhTable += buatTablePerJurusan(folderDBHariIni, daftarJurusan[idxJurusan].Name(), fmt.Sprintf("%v, %v %v %v", hariBesok, hariBesokTgl, hariBesokBln, hariBesokThn))
			} else if i == 2 {
				templateSeluruhTable += buatTablePerJurusan(folderDBHariIni, daftarJurusan[idxJurusan].Name(), fmt.Sprintf("%v, %v %v %v", hariLusa, hariLusaTgl, hariLusaBln, hariLusaThn))
			}

			templateSeluruhTable += "<br>"
		}
	}

	templateSeluruhHalaman = strings.ReplaceAll(antrian, "(REPLACE-INI-DENGAN-TABLE)", templateSeluruhTable)

	return templateSeluruhHalaman
}
