package views

import (
	"bufio"
	"fmt"
	"os"
	"project_linkedlist_uts_2026/models"
	"strconv"
	"strings"
)

func PaketInsertView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===== TAMBAH PAKESET =====")

	fmt.Print("Nama Paket : ")
	namaPaket, _ := reader.ReadString('\n')
	namaPaket = strings.TrimSpace(namaPaket)

	fmt.Print("Harga : ")
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	harga, _ := strconv.Atoi(hargaStr)

	fmt.Print("Keterangan : ")
	keterangan, _ := reader.ReadString('\n')
	keterangan = strings.TrimSpace(keterangan)

	fmt.Println(" list menu ")
	MenuViewAll()
	fmt.Println("---------------------------")

	fmt.Print("ID Menu Pembuka : ")
	idMenuPembukaStr, _ := reader.ReadString('\n')
	idMenuPembukaStr = strings.TrimSpace(idMenuPembukaStr)
	idMenuPembuka, _ := strconv.Atoi(idMenuPembukaStr)

	fmt.Print("ID Menu Utama : ")
	idMenuUtamaStr, _ := reader.ReadString('\n')
	idMenuUtamaStr = strings.TrimSpace(idMenuUtamaStr)
	idMenuUtama, _ := strconv.Atoi(idMenuUtamaStr)

	fmt.Print("ID Menu Penutup : ")
	idMenuPenutupStr, _ := reader.ReadString('\n')
	idMenuPenutupStr = strings.TrimSpace(idMenuPenutupStr)
	idMenuPenutup, _ := strconv.Atoi(idMenuPenutupStr)

	status := models.PaketInsert(namaPaket, idMenuPembuka, idMenuUtama, idMenuPenutup, harga)

	if status {
		fmt.Println("Paket berhasil ditambahkan")
	} else {
		fmt.Println("Gagal menambahkan paket")
	}
}

func PaketViewAll() {
	data := models.PaketView()
	if len(data) == 0 {
		fmt.Println("Data paket masih kosong")
		return
	}
	fmt.Println("===== DAFTAR PAKESET =====")
	for _, p := range data {
		fmt.Println("Nama Paket :", p.NamaPaket)
		fmt.Println("Harga      :", p.Harga)
		fmt.Println("---------------------------")
		fmt.Println("Menu Pembuka :", models.GetNamaMenuById(p.IdMenuPembuka))
		fmt.Println("Menu Utama   :", models.GetNamaMenuById(p.IdMenuUtama))
		fmt.Println("Menu Penutup :", models.GetNamaMenuById(p.IdMenuPenutup))
		fmt.Println("---------------------------")
		fmt.Println()
	}
}

func PaketUpdateView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===== UPDATE PAKESET =====")

	fmt.Print("ID Paket : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Print("Nama Baru : ")
	namaPaket, _ := reader.ReadString('\n')
	namaPaket = strings.TrimSpace(namaPaket)

	fmt.Print("Harga Baru : ")
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	harga, _ := strconv.Atoi(hargaStr)

	fmt.Println(" list menu ")
	MenuViewAll()
	fmt.Println("---------------------------")

	fmt.Print("ID Menu Pembuka : ")
	idMenuPembukaStr, _ := reader.ReadString('\n')
	idMenuPembukaStr = strings.TrimSpace(idMenuPembukaStr)
	idMenuPembuka, _ := strconv.Atoi(idMenuPembukaStr)

	fmt.Print("ID Menu Utama : ")
	idMenuUtamaStr, _ := reader.ReadString('\n')
	idMenuUtamaStr = strings.TrimSpace(idMenuUtamaStr)
	idMenuUtama, _ := strconv.Atoi(idMenuUtamaStr)

	fmt.Print("ID Menu Penutup : ")
	idMenuPenutupStr, _ := reader.ReadString('\n')
	idMenuPenutupStr = strings.TrimSpace(idMenuPenutupStr)
	idMenuPenutup, _ := strconv.Atoi(idMenuPenutupStr)

	models.PaketUpdate(id, namaPaket, harga, idMenuPembuka, idMenuUtama, idMenuPenutup)
}

func PaketDeleteView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===== HAPUS PAKESET =====")

	fmt.Print("ID Paket : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	hasil := models.PaketDelete(id)

	if hasil {
		fmt.Println("Paket berhasil dihapus")
	} else {
		fmt.Println("Gagal menghapus paket")
	}
}
