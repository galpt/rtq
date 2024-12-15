package handlers

/*
Berdasarkan teks yang diberikan, jam operasional SASC BINUS UNIVERSITY adalah:

Senin - Kamis: 09.00 - 16.00 WIB, istirahat 12.00 - 13.00 WIB
Jumat: 09.00 - 16.00 WIB, istirahat 11.30 - 13.00 WIB
Sabtu: 09.00 - 14.00 WIB, istirahat 12.00 - 13.00 WIB
*/

type (
	Mahasiswa struct {
		Nim          string `form:"nim"`
		Nama         string `form:"nama"`
		Jurusan      string `form:"jurusan"`
		JenisKonsul  string `form:"jeniskonsul"` // Akademik atau Non-Akademik
		DurasiKonsul string // format HH:MM-HH:MM, contoh "10:00-10:45"
	}

	Antrian struct {
		NoAntrian     string
		Nim           string
		SudahDilayani bool
		WaktuAntri    string // timestamp
	}

	Konselor struct {
		IDKonselor   string `form:"idkonselor"`
		NamaKonselor string `form:"namakonselor"`
	}
)
