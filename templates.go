// templates.go
package main

import (
	"time"
)

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
    <title>Beranda</title>
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
                    class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
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

    <script>
        function openModal(modalId) {
			const nimInput = document.getElementById('nim');
            const nim = nimInput.value.trim();
			const nameInput = document.getElementById('name');
            const nama = nameInput.value.trim();
            const jurusan = document.getElementById('major').value;
            const jamKonsul = document.getElementById('jam_konsultasi').value;
            const jenisKonsul = document.getElementById('jenis_konsultasi').value;
			
			const errorMessages = document.querySelectorAll('.error-message');
            errorMessages.forEach(msg => msg.remove());
            nimInput.classList.remove('border-red-500');
			nameInput.classList.remove('border-red-500');

            let isValid = true;

            if (nim === '') {
                displayError(nimInput, 'NIM wajib diisi');
                isValid = false;
            } else if (!/^\d+$/.test(nim)) {
                displayError(nimInput, 'NIM harus berupa angka');
                isValid = false;
            }
			
			if (nama === '') {
                displayError(nameInput, 'Nama wajib diisi');
                isValid = false;
            }
			
			if (!isValid) {
				return;
			}

            document.getElementById('modalNim').textContent = nim;
            document.getElementById('modalNama').textContent = nama;
            document.getElementById('modalJurusan').textContent = jurusan;
            document.getElementById('modalJamKonsul').textContent = formatJamKonsul(jamKonsul);
            document.getElementById('modalJenisKonsul').textContent = capitalizeFirstLetter(jenisKonsul);

            document.getElementById(modalId).style.display = 'block';
            document.getElementsByTagName('body')[0].classList.add('overflow-y-hidden');
        }

        function closeModal(modalId) {
            document.getElementById(modalId).style.display = 'none';
            document.getElementsByTagName('body')[0].classList.remove('overflow-y-hidden');
        }
	
	    function submitForm() {
		    document.getElementById('action').value = 'konfirmasi';
		    document.getElementById('daftarForm').submit();
	    }
	
		function formatJamKonsul(jamKonsul) {
			const parts = jamKonsul.split('|');
			if (parts.length === 2) {
				const hari = parts[0];
				const jam = parts[1];
				
				let hariFormatted = "";
				if (hari === "hariini") {
					hariFormatted = "Hari Ini";
				} else if (hari === "besok") {
					hariFormatted = "Besok";
				} else if (hari === "lusa") {
					hariFormatted = "Lusa";
				}
				
				return hariFormatted + ", " + jam;
			}
			return jamKonsul;
		}
		
		function capitalizeFirstLetter(string) {
			return string.charAt(0).toUpperCase() + string.slice(1);
		}
	
        // Close all modals when press ESC
        document.onkeydown = function(event) {
            event = event || window.event;
            if (event.keyCode === 27) {
                document.getElementsByTagName('body')[0].classList.remove('overflow-y-hidden');
                let modals = document.getElementsByClassName('modal');
                Array.prototype.slice.call(modals).forEach(i => {
                    i.style.display = 'none';
                });
            }
        };

        function validateForm() {
            const nimInput = document.getElementById('nim');
            const nim = nimInput.value.trim();
			const nameInput = document.getElementById('name');
            const name = nameInput.value.trim();
            const errorMessages = document.querySelectorAll('.error-message');
            errorMessages.forEach(msg => msg.remove());
            nimInput.classList.remove('border-red-500');
			nameInput.classList.remove('border-red-500');

            let isValid = true;

            if (nim === '') {
                displayError(nimInput, 'NIM wajib diisi');
                isValid = false;
            } else if (!/^\d+$/.test(nim)) {
                displayError(nimInput, 'NIM harus berupa angka');
                isValid = false;
            }
			
			if (name === '') {
                displayError(nameInput, 'Nama wajib diisi');
                isValid = false;
            }

            return isValid;
        }

        function displayError(inputField, message) {
            const errorMessage = document.createElement('p');
            errorMessage.textContent = message;
            errorMessage.classList.add('text-red-500', 'text-sm', 'mt-1', 'error-message');
            inputField.parentNode.insertBefore(errorMessage, inputField.nextSibling);
            inputField.classList.add('border-red-500');
        }

        function submitForm() {
            if (validateForm()) {
                document.getElementById('action').value = 'konfirmasi';
                document.getElementById('daftarForm').submit();
            }
        }
    </script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="max-w-xl mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Daftar Konsultasi</h1>
        </div>

        <form id="daftarForm" class="space-y-6" action="/daftar" method="post">
            
            (REPLACE-INI-DENGAN-FORM-INPUT)
			
			<div class="mt-6">
				<button type="button" onclick="openModal('modelConfirm')"
					class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
					Daftar
				</button>
			</div>
			
			<input type="hidden" id="action" name="action" value="">
        </form>
		
		<div id="modelConfirm" class="fixed hidden z-50 inset-0 bg-gray-900 bg-opacity-60 overflow-y-auto h-full w-full px-4 ">
		<div class="relative top-40 mx-auto shadow-xl rounded-md bg-white max-w-md">
			<div class="flex justify-end p-2">
				<button onclick="closeModal('modelConfirm')" type="button"
					class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center">
					<svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
						<path fill-rule="evenodd"
							d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
							clip-rule="evenodd"></path>
					</svg>
				</button>
			</div>
			<div class="p-6 pt-0 text-center">
				<svg class="w-20 h-20 text-yellow-600 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
						d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
				</svg>
				<h3 class="text-xl font-normal text-gray-500 mt-5 mb-6">Apakah data berikut sudah benar?</h3>
				
				<div class="text-left space-y-2 mb-4">
					<p><span class="font-semibold">NIM:</span> <span id="modalNim"></span></p>
					<p><span class="font-semibold">Nama:</span> <span id="modalNama"></span></p>
					<p><span class="font-semibold">Jurusan:</span> <span id="modalJurusan"></span></p>
					<p><span class="font-semibold">Jam Konsultasi:</span> <span id="modalJamKonsul"></span></p>
					<p><span class="font-semibold">Jenis Konsultasi:</span> <span id="modalJenisKonsul"></span></p>
				</div>

				<button onclick="submitForm()"
					class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-base inline-flex items-center px-3 py-2.5 text-center mr-2">
					Ya, daftarkan saya
				</button>
				<button onclick="closeModal('modelConfirm')"
					class="text-gray-900 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-cyan-200 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center"
					data-modal-toggle="delete-user-modal">
					Ubah data
				</button>
			</div>
		</div>
	</div>`

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

	modalKonfirmasi = `<div id="modelConfirm" class="fixed hidden z-50 inset-0 bg-gray-900 bg-opacity-60 overflow-y-auto h-full w-full px-4 ">
		<div class="relative top-40 mx-auto shadow-xl rounded-md bg-white max-w-md">
			<div class="flex justify-end p-2">
				<button onclick="closeModal('modelConfirm')" type="button"
					class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center">
					<svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
						<path fill-rule="evenodd"
							d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
							clip-rule="evenodd"></path>
					</svg>
				</button>
			</div>
			<div class="p-6 pt-0 text-center">
				<svg class="w-20 h-20 text-yellow-600 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
						d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
				</svg>
				<h3 class="text-xl font-normal text-gray-500 mt-5 mb-6">Apakah data berikut sudah benar?</h3>
				
				<div class="text-left space-y-2 mb-4">
					<p><span class="font-semibold">NIM:</span> <span id="modalNim"></span></p>
					<p><span class="font-semibold">Nama:</span> <span id="modalNama"></span></p>
					<p><span class="font-semibold">Jurusan:</span> <span id="modalJurusan"></span></p>
					<p><span class="font-semibold">Jam Konsultasi:</span> <span id="modalJamKonsul"></span></p>
					<p><span class="font-semibold">Jenis Konsultasi:</span> <span id="modalJenisKonsul"></span></p>
				</div>

				<button onclick="submitForm()"
					class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-base inline-flex items-center px-3 py-2.5 text-center mr-2">
					Ya, daftarkan saya
				</button>
				<button onclick="closeModal('modelConfirm')"
					class="text-gray-900 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-cyan-200 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center"
					data-modal-toggle="delete-user-modal">
					Ubah data
				</button>
			</div>
		</div>
	</div>`

	loginAdmin = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Masuk Admin</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="max-w-xl mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Masuk Admin</h1>
        </div>

        <!-- Login Form -->
        <form class="space-y-6" action="/admin/login" method="post">
            <!-- Username Input -->
            <div class="space-y-2">
                <label for="username" class="block text-sm font-medium text-gray-700">
                    Username
                </label>
                <input type="text" id="username" name="username"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan username" required>
            </div>

            <!-- Password Input -->
            <div class="space-y-2">
                <label for="password" class="block text-sm font-medium text-gray-700">
                    Password
                </label>
                <input type="password" id="password" name="password"
                    class="w-full rounded-lg border-2 border-gray-200 p-3 text-gray-700 focus:border-blue-500 focus:outline-none transition duration-200"
                    placeholder="Masukkan password" required>
            </div>

            <!-- Submit Button -->
            <div class="mt-6">
                <button type="submit"
                    class="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 transform hover:-translate-y-1 hover:shadow-lg">
                    Masuk
                </button>
            </div>
			
			<div class="mt-4 text-sm text-center">
				<a href="#" class="text-blue-500 hover:underline">Lupa Password?</a>
			</div>
        </form>
    </div>
</body>

</html>`

	dashboardAdmin = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Dasbor Admin</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="min-h-screen bg-gradient-to-br from-blue-400 to-blue-600 p-4">
    <div class="max-w-xl mx-auto mt-10 bg-white/95 rounded-xl shadow-xl p-6 md:p-8">
        <!-- Header -->
        <div class="bg-blue-800 text-white p-4 rounded-lg mb-8 text-center">
            <h1 class="text-2xl font-bold">Dasbor Admin</h1>
        </div>

        <!-- Navigation -->
        <nav class="mb-8">
            <ul class="flex space-x-4 justify-center">
                <li><a href="/admin/dashboard" class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg transition duration-200">Dasbor</a></li>
                <li><a href="/antrian" class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold py-2 px-4 rounded-lg transition duration-200">Daftar Antrian</a></li>
            </ul>
        </nav>

        <!-- Content Area -->
        <div class="bg-gray-100 p-4 rounded-lg">
            <p>Selamat datang di dasbor admin!</p>
        </div>
    </div>
</body>

</html>`
)
