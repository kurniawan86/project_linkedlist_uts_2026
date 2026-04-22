package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	controller "project_linkedlist_uts_2026/controllers"
)

func UserInsertView() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===== TAMBAH USER =====")

	fmt.Print("Nama User : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Role [admin/kasir/pelanggan] : ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(strings.ToLower(role))

	status, err := controller.UserInsertController(nama, "", role)
	if err != nil {
		fmt.Println("Gagal menambahkan user:", err.Error())
		return
	}

	if status {
		fmt.Println("User berhasil ditambahkan")
		return
	}

	fmt.Println("User gagal ditambahkan")
}

func UserViewAll() {
	dataUsers := controller.ShowAllUserController()
	if len(dataUsers) == 0 {
		fmt.Println("Data user masih kosong")
		return
	}

	fmt.Println("===== DAFTAR USER =====")
	for _, user := range dataUsers {
		fmt.Println("ID User      :", user.IdUser)
		fmt.Println("Nama User    :", user.NamaUser)
		fmt.Println("Role         :", user.Role)
		fmt.Println("Date Created :", user.DateCreated)
		fmt.Println("Date Updated :", user.DateUpdated)
		fmt.Println("---------------------------")
		fmt.Println()
	}
}
