package main

import "time"

// bagian ini untuk semua data yang bersifat template

const (
	timeFormat             = time.RFC3339
	namaFilePengaturan     = "./pengaturan.txt"
	namaFileDataMahasiswa  = "data_mahasiswa.txt"
	namaFileWaktuReservasi = "waktu_reservasi.txt"
	namaFileJenisKonsul    = "jenis_konsul.txt"
	folderDB               = "./db"
	folderDBHariIni        = "./db/hariini"
	folderDBBesok          = "./db/besok"
	folderDBLusa           = "./db/lusa"
	totalHari              = 3
)

const (

	// =========================
	// bagian ini untuk semua template pengaturan nya
	// =========================

	templateDefaultPengaturan = `// harus dipisah dengan tanda '|'
jurusan=[Computer Science|Information Systems|International Relations]

// nama konselor harus berurutan dengan posisi 'jurusan' yang diisi di atas
namakonselor=[Ita|Ika|Agus]

// jam konsul sesuaikan dengan format seperti ini
jamkonsul=[09.00-09.45|09.45-10.30|10.30-11.15|11.15-12.00|13.00-13.45|13.45-14.30|14.30-15.15|15.15-16.00]`

	templateDefaultDataMahasiswa = `nimnama=[1234567890|Galih]`

	// =========================
	// bagian ini untuk semua template html nya
	// =========================

	home = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="max-w-xl mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Sistem Antrian</h1>
        </div>

        <!-- Buttons -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-8">

            <a href="http://localhost:8080/antrian">
                <button type="button"
                    class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Cek Daftar Antrian
                </button>
            </a>

            <a href="http://localhost:8080/daftar">
                <button type="button"
                    class="w-full bg-green-500 hover:bg-green-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Daftar Konsultasi
                </button>
            </a>

        </div>
    </div>
</body>

</html>`

	daftar = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Daftar Konsultasi</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="max-w-xl mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Daftar Konsultasi</h1>
        </div>

        <form class="space-y-6" action="/daftar" method="post">
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

            (REPLACE-INI-DENGAN-DROPDOWN-MENU)

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

            <!-- Tombol daftar -->
            <div class="mt-8">
                <button type="submit" formmethod="post" type="button"
                    class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Daftar
                </button>
            </div>
            
        </form>
    </div>
</body>

</html>`

	antrian = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Daftar Antrian</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Daftar Antrian</h1>
        </div>

        (REPLACE-INI-DENGAN-TABLE)


    </div>
</body>

</html>`

	alertSuksesDaftarKonsul = `<div id="alert-3"
            class="flex items-center p-4 mb-4 text-green-800 rounded-lg bg-green-50 dark:bg-gray-800 dark:text-green-400"
            role="alert">
            <svg class="flex-shrink-0 w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                viewBox="0 0 20 20">
                <path
                    d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
            </svg>
            <span class="sr-only">Info</span>
            <div class="ms-3 text-sm font-medium">
                Berhasil mendaftarkan pada sesi yang diinginkan.
            </div>
        </div>`

	alertGagalDaftarKonsul = `<div id="alert-2"
            class="flex items-center p-4 mb-4 text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
            role="alert">
            <svg class="flex-shrink-0 w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                viewBox="0 0 20 20">
                <path
                    d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
            </svg>
            <span class="sr-only">Info</span>
            <div class="ms-3 text-sm font-medium">
                Cek kembali daftar antrian, mungkin sesi tersebut sudah direservasi oleh
                orang lain. <a href="/daftar" class="font-semibold underline hover:no-underline">Kembali ke halaman
                    daftar konsultasi</a>.
            </div>
        </div>`

	alertMenyiapkanDaftarAntrian = `<div id="alert-4"
            class="flex items-center p-4 mb-4 text-yellow-800 rounded-lg bg-yellow-50 dark:bg-gray-800 dark:text-yellow-300"
            role="alert">
            <svg class="flex-shrink-0 w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                viewBox="0 0 20 20">
                <path
                    d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
            </svg>
            <span class="sr-only">Info</span>
            <div class="ms-3 text-sm font-medium">
                Daftar antrian sedang disiapkan. Mohon refresh halaman ini setelah beberapa detik. <a href="/antrian"
                    class="font-semibold underline hover:no-underline">Refresh kembali</a>.
            </div>
        </div>`
)
