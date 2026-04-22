package domain

type TransaksiPenjualan struct {
	IdPenjualan      int
	TanggalPenjualan string
	JenisTransaksi   string
	TotalTransaksi   int
	Diskon           float32
	DP               int
	Detail           [10]DetailTransaksi
	JumlahPaket      int
}

type DetailTransaksi struct {
	Paket    PaketMenu
	Jumlah   int
	Subtotal int
}
