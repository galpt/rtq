# duarr-kopkar v2

> [!NOTE]
>
> 1. Source code ini bersifat sebagai kerangka atau fondasi dan dapat di-custom sesuai kebutuhan.
>
> 2. Sejak program ini dibuat hingga saat README ini ditulis, program telah diuji pada sistem operasi Windows saja. Seharusnya tidak ada masalah jika program dijalankan di server/cloud dengan sistem operasi berbasis Linux.

## Cara menggunakan

1. Download & install bahasa pemrograman Go (https://go.dev/doc/install) sesuai dengan sistem operasi kalian.
2. Pastikan Go sudah terinstall dengan benar dengan mengetik `go version` pada Command Prompt setelah selesai melakukan instalasi, seperti berikut:

```yaml
C:\Users\galpt>go version
go version go1.22.4 windows/amd64
```

3. Download repository ini, dan extract ZIP-nya di suatu folder kosong yang Anda inginkan.
4. Pada folder tersebut, buka Command Prompt dan ketik `go mod tidy` untuk memastikan program berfungsi dengan baik.
5. Gunakan `go build` untuk mengubah source code menjadi file `.exe` atau gunakan `go run main.go` jika ingin langsung mencoba program tanpa proses compile.

Berikut adalah tampilan saat program berhasil berjalan:

```yaml
E:\GitHub\duarr-kopkar-v2>go run main.go
Server started.
Access on:
 - http://0.0.0.0:8080/
 - http://127.0.0.1:8080/
```

> [!IMPORTANT]
>
> Jangan tutup Command Prompt atau terminal Anda jika ingin progran tetap berjalan karena ini akan memaksa program untuk berhenti.
