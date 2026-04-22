package views

import (
	"bufio"
	"fmt"
	"os"
	"project_linkedlist_uts_2026/domain"
	model "project_linkedlist_uts_2026/models"
	"strings"
)

func TransaksiInsertView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Id Penjualan : ")
	var idPenjualan int
	fmt.Scan(&idPenjualan)

	fmt.Print("Masukkan Jenis Transaksi : ")
	var jenisTransaksi string
	fmt.Scan(&jenisTransaksi)

	fmt.Print("Masukkan DP : ")
	var dp int
	fmt.Scan(&dp)

	var jumlahPaket int = 0
	var detail [10]domain.DetailTransaksi

	// menampilkan paket
	PaketViewAll()

	for {
		fmt.Print("Nama Paket : ")
		namaPaket, _ := reader.ReadString('\n')
		namaPaket = strings.TrimSpace(namaPaket)

		detail[jumlahPaket].Paket = model.GetPaketByName(namaPaket)
		fmt.Print("Masukkan Jumlah : ")
		var jumlah int
		fmt.Scanln(&jumlah)
		detail[jumlahPaket].Jumlah = jumlah

		//code untuk menghitung subtotal per paket
		detail[jumlahPaket].Subtotal = detail[jumlahPaket].Paket.Harga * detail[jumlahPaket].Jumlah

		jumlahPaket++

		fmt.Print("Apakah ingin menambahkan paket lagi? (y/n)")
		var tambah string
		fmt.Scanln(&tambah)
		if tambah == "n" {
			break
		}
	}

	totalTransaksi := 0
	for i := 0; i < jumlahPaket; i++ {
		totalTransaksi = totalTransaksi + detail[i].Subtotal
	}

	if model.TransaksiInsert(idPenjualan, jenisTransaksi, totalTransaksi, dp, jumlahPaket, detail) {
		fmt.Println("Transaksi berhasil")
	} else {
		fmt.Println("Transaksi gagal")
	}
}

func TransaksiViewAll() {
	for _, transaksi := range model.TransaksiView() {
		fmt.Println("Id Penjualan : ", transaksi.IdPenjualan)
		fmt.Println("Tanggal Penjualan : ", transaksi.TanggalPenjualan)
		fmt.Println("Jenis Transaksi : ", transaksi.JenisTransaksi)
		fmt.Println("DP : ", transaksi.DP)
		fmt.Println("Jumlah Paket : ", transaksi.JumlahPaket)
		fmt.Println("====================================")
		for i := 0; i < transaksi.JumlahPaket; i++ {
			fmt.Println("Nama Paket : ", transaksi.Detail[i].Paket.NamaPaket)
			fmt.Println("Jumlah : ", transaksi.Detail[i].Jumlah)
			fmt.Println("Subtotal : ", transaksi.Detail[i].Subtotal)
			fmt.Println("====================================")
		}
		fmt.Println("Total Transaksi : ", transaksi.TotalTransaksi)
		fmt.Println("Sisa Transaksi : ", transaksi.TotalTransaksi-transaksi.DP)
		fmt.Println("====================================")
	}
}
