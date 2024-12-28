## Tentang

RTQ adalah sistem manajemen antrian yang dirancang agar dapat memberikan *reliability* tinggi dalam mengatur antrian secara *real-time*. RTQ dibuat sebagai upaya untuk mengatasi masalah ketidakefisienan di mana instansi pendidikan yang memiliki layanan konsultasi mahasiswa masih menggunakan cara manual untuk melayani antrian konsultasi, mulai dari ketika mahasiswa ingin mendaftarkan sesi konsultasi hingga proses konsultasi selesai.

Tidak adanya sebuah sistem yang memberikan informasi terbaru secara *real-time* kepada mahasiswa seputar berapa panjang antrian saat ini, apakah masih dapat mendaftar konsultasi untuk hari itu atau tidak, kejelasan durasi konsultasi, dan masalah sejenis membuat proses konsultasi menjadi sulit karena ada faktor yang tidak bisa diprediksi (misal antrian hari itu sudah penuh atau masih dapat konsultasi), dan itu akan menyebabkan masalah lain (misal sulit memprediksi waktu perjalanan mahasiswa dari rumah ke kampus hanya untuk konsultasi, tapi ternyata sudah tidak bisa konsultasi karena antrian hari itu sudah penuh).

Fungsi RTQ adalah memastikan proses konsultasi menjadi lebih terprediksi agar dapat meminimalisir kerugian waktu, tenaga, dan biaya transportasi dengan menjamin apakah mahasiswa dapat melakukan konsultasi atau tidak pada waktu yang diinginkan.

## Cara menggunakan

1. Download & install bahasa pemrograman Go (https://go.dev/doc/install) sesuai dengan sistem operasi kalian.
2. Pastikan Go sudah terinstall dengan benar dengan mengetik `go version` pada Command Prompt setelah selesai melakukan instalasi, seperti berikut:

```yaml
C:\Users\galpt>go version
go version go1.23.4 windows/amd64
```

3. Download repository ini, dan extract ZIP-nya di suatu folder kosong yang Anda inginkan.
4. Pada folder tersebut, buka Command Prompt dan ketik `go mod tidy` untuk memastikan program berfungsi dengan baik.
5. Gunakan `go build` untuk mengubah source code menjadi file `.exe` atau gunakan `go run main.go` jika ingin langsung mencoba program tanpa proses compile.

Berikut adalah tampilan saat program berhasil berjalan:

```yaml
D:\Disk D\GitHub\rtq>go run main.go
Server jalan di 0.0.0.0:8080
```

> [!IMPORTANT]
>
> Jangan tutup Command Prompt atau terminal Anda jika ingin progran tetap berjalan karena ini akan memaksa program untuk berhenti.
