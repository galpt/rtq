package main

// bagian ini untuk semua data yang bersifat template

const (
	namaFilePengaturan    = "./pengaturan.txt"
	namaFileDataMahasiswa = "data_mahasiswa.txt"
	folderDB              = "./db"
	folderDBHariIni       = "./db/hariini"
	folderDBBesok         = "./db/besok"
	folderDBLusa          = "./db/lusa"
	totalHari             = 3
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

        <form class="space-y-6">
            <!-- NIM Input -->
            <div class="space-y-2">
                <label for="nim" class="block text-sm font-medium text-gray-700">
                    NIM (Nomor Induk Mahasiswa)
                </label>
                <input type="text" id="nim"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan NIM" required>
            </div>

            <!-- Name Input -->
            <div class="space-y-2">
                <label for="name" class="block text-sm font-medium text-gray-700">
                    Nama
                </label>
                <input type="text" id="name"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan nama lengkap" required>
            </div>

            <!-- Major Select -->
            <div class="space-y-2">
                <label for="major" class="block text-sm font-medium text-gray-700">
                    Jurusan
                </label>
                <select id="major"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    required>
                    <option value="" disabled>Pilih jurusan</option>
                    <option value="cs">Computer Science</option>
                    <option value="is">System Information</option>
                    <option value="ir">International Relations</option>
                </select>
            </div>

            <!-- Pilih jam konsul -->
            <div class="space-y-2">
                <label for="major" class="block text-sm font-medium text-gray-700">
                    Jam Konsultasi
                </label>
                <select id="major"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    required>
                    <option value="" disabled>Pilih jam konsultasi</option>
                    <option value="09.00-09.45">09.00-09.45</option>
                    <option value="09.45-10.30">09.45-10.30</option>
                    <option value="10.30-11.15">10.30-11.15</option>
                </select>
            </div>

            <!-- Buttons -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-8">
                <button type="button"
                    class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Konsultasi Akademik
                </button>
                <button type="button"
                    class="w-full bg-green-500 hover:bg-green-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Konsultasi Non-Akademik
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
)
