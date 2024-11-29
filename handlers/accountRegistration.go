package handlers

type (
	Account struct {
		Nama        string `form:"nama"`
		Email       string `form:"email"`
		Username    string `form:"username"`
		Password    string `form:"password"`
		NoHandphone int    `form:"noHp"`
	}

	Pinjaman struct {
		Saldo             float64  `json:"saldo"`
		RiwayatPeminjaman []string `json:"riwayatPinjaman"`
		JenisPinjaman     []int    `json:"jenisPinjaman"`
	}
)
