package models

import (
	"errors"
	"project_linkedlist_uts_2026/domain"
	"time"
)

var tabelTransaksi [100]domain.TransaksiPenjualan
var indexTransaksi int = 0

func TransaksiInsert(idPenjualan int, JenisTransaksi string,
	totalTransaksi int, dp int, jumlahPaket int, detail [10]domain.DetailTransaksi) bool {

	tabelTransaksi[indexTransaksi].IdPenjualan = idPenjualan
	tabelTransaksi[indexTransaksi].TanggalPenjualan = time.Now().Format("2006-01-02")
	tabelTransaksi[indexTransaksi].JenisTransaksi = JenisTransaksi
	tabelTransaksi[indexTransaksi].TotalTransaksi = totalTransaksi
	tabelTransaksi[indexTransaksi].DP = dp
	tabelTransaksi[indexTransaksi].JumlahPaket = jumlahPaket
	tabelTransaksi[indexTransaksi].Detail = detail
	indexTransaksi++
	return true
}

func TransaksiView() []domain.TransaksiPenjualan {
	return tabelTransaksi[:indexTransaksi]
}

func TransaksiUpdate(id int, JenisTransaksi string, totalTransaksi int, dp int,
	jumlahPaket int, detail [10]domain.DetailTransaksi) (domain.TransaksiPenjualan, error) {

	if id < 0 || id >= indexTransaksi {
		return domain.TransaksiPenjualan{}, errors.New("id tidak ditemukan")
	} else {
		if JenisTransaksi != "" {
			tabelTransaksi[id].JenisTransaksi = JenisTransaksi
		}
		if totalTransaksi != 0 {
			tabelTransaksi[id].TotalTransaksi = totalTransaksi
		}
		if dp != 0 {
			tabelTransaksi[id].DP = dp
		}
		if jumlahPaket != 0 {
			tabelTransaksi[id].JumlahPaket = jumlahPaket
		}
		if detail != [10]domain.DetailTransaksi{} {
			tabelTransaksi[id].Detail = detail
		}
		return tabelTransaksi[id], nil
	}
}

func TransaksiDelete(id int) bool {
	if id < 0 || id >= indexTransaksi {
		return false
	} else {
		tabelTransaksi[id] = tabelTransaksi[indexTransaksi-1]
		indexTransaksi--
		return true
	}
}
