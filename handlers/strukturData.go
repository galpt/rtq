package handlers

type (
	Account struct {
		Nim         string `form:"nim"`
		Nama        string `form:"nama"`
		Jurusan     string `form:"jurusan"`
		WaktuKonsul string `form:"waktukonsul"`
	}
)
