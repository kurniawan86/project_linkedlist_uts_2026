package models

import "project_linkedlist_uts_2026/domain"

var tabelPaket [100]domain.PaketMenu
var indexPaket int = 0

func PaketInsert(nama string, idMenuPembuka int, idMenuUtama int, idMenuPenutup int, harga int) bool {
	if nama == "" || harga == 0 {
		return false
	}
	tabelPaket[indexPaket].NamaPaket = nama
	tabelPaket[indexPaket].IdMenuPembuka = idMenuPembuka
	tabelPaket[indexPaket].IdMenuUtama = idMenuUtama
	tabelPaket[indexPaket].IdMenuPenutup = idMenuPenutup
	tabelPaket[indexPaket].Harga = harga
	indexPaket++
	return true
}

func PaketView() []domain.PaketMenu {
	return tabelPaket[:indexPaket]
}

func PaketUpdate(id int, nama string, idMenuPembuka int, idMenuUtama int, idMenuPenutup int, harga int) domain.PaketMenu {

	if id < 0 || id >= indexPaket {
		return domain.PaketMenu{}
	}

	if nama != "" {
		tabelPaket[id].NamaPaket = nama
	}

	if harga != 0 {
		tabelPaket[id].Harga = harga
	}

	tabelPaket[id].IdMenuPembuka = idMenuPembuka
	tabelPaket[id].IdMenuUtama = idMenuUtama
	tabelPaket[id].IdMenuPenutup = idMenuPenutup

	return tabelPaket[id]
}

func PaketDelete(id int) bool {

	if id < 0 || id >= indexPaket {
		return false
	}

	tabelPaket[id] = tabelPaket[indexPaket-1]
	indexPaket--
	return true
}

func PaketInsertStatic() {
	PaketInsert("paket keluarga", 0, 1, 2, 100000)
	PaketInsert("paket pacaran", 5, 3, 2, 150000)
	PaketInsert("paket hemat", 4, 3, 1, 300000)
}

func GetPaketByName(nama string) domain.PaketMenu {
	for i := 0; i < indexPaket; i++ {
		if tabelPaket[i].NamaPaket == nama {
			return tabelPaket[i]
		}
	}
	return domain.PaketMenu{}
}
