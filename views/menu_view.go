package views

import (
	"bufio"
	"fmt"
	"os"
	"project_linkedlist_uts_2026/models"
	"strconv"
	"strings"
)

func MenuInsertView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===== TAMBAH MENU =====")

	fmt.Print("Nama Menu : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Harga : ")
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	harga, _ := strconv.Atoi(hargaStr)

	fmt.Print("Keterangan : ")
	ket, _ := reader.ReadString('\n')
	ket = strings.TrimSpace(ket)

	status := models.MenuInsert(nama, harga, ket)

	if status {
		fmt.Println("Menu berhasil ditambahkan")
	} else {
		fmt.Println("Gagal menambahkan menu")
	}
}

func MenuViewAll() {

	data := models.MenuView()
	if len(data) == 0 {
		fmt.Println("Data menu masih kosong")
		return
	}
	fmt.Println("===== DAFTAR MENU =====")
	for _, m := range data {
		fmt.Println("ID         :", m.IdMenu)
		fmt.Println("Nama Menu  :", m.NamaMenu)
		fmt.Println("Harga      :", m.Harga)
		fmt.Println("Keterangan :", m.Keterangan)
		fmt.Println("---------------------------")
		fmt.Println()
	}
}

func MenuUpdateView() {

	var id int
	var nama string
	var harga int
	var ket string

	fmt.Println("===== UPDATE MENU =====")

	fmt.Print("ID Menu : ")
	fmt.Scanln(&id)

	fmt.Print("Nama Baru : ")
	fmt.Scanln(&nama)

	fmt.Print("Harga Baru : ")
	fmt.Scanln(&harga)

	fmt.Print("Keterangan Baru : ")
	fmt.Scanln(&ket)

	data := models.MenuUpdate(id, nama, harga, ket)

	if data.NamaMenu == "" {
		fmt.Println("Data menu tidak ditemukan")
		return
	}

	fmt.Println("Menu berhasil diupdate")
	fmt.Println("ID         :", data.IdMenu)
	fmt.Println("Nama Menu  :", data.NamaMenu)
	fmt.Println("Harga      :", data.Harga)
	fmt.Println("Keterangan :", data.Keterangan)
}

func MenuDeleteView() {

	var id int

	fmt.Println("===== HAPUS MENU =====")

	fmt.Print("ID Menu : ")
	fmt.Scanln(&id)

	status := models.MenuDelete(id)

	if status {
		fmt.Println("Menu berhasil dihapus")
	} else {
		fmt.Println("Menu tidak ditemukan")
	}
}
