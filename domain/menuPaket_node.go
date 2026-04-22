package domain

type Menu struct {
	IdMenu     int
	NamaMenu   string
	Harga      int
	Keterangan string
}

type PaketMenu struct {
	NamaPaket     string
	IdMenuPembuka int
	IdMenuUtama   int
	IdMenuPenutup int
	Harga         int
}
