package handlers

/*
Berdasarkan teks yang diberikan, jam operasional SASC BINUS UNIVERSITY adalah:

Senin - Kamis: 09.00 - 16.00 WIB, istirahat 12.00 - 13.00 WIB
Jumat: 09.00 - 16.00 WIB, istirahat 11.30 - 13.00 WIB
Sabtu: 09.00 - 14.00 WIB, istirahat 12.00 - 13.00 WIB
*/

type (
	Mahasiswa struct {
		Nim         string `form:"nim"`
		Nama        string `form:"nama"`
		Jurusan     string `form:"jurusan"`
		WaktuKonsul string `form:"waktukonsul"`
	}
)
