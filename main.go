package main

import (
	"fmt"
	"net/http"
	controller "project_linkedlist_uts_2026/controllers"
	"project_linkedlist_uts_2026/handlers"
	view "project_linkedlist_uts_2026/views"
)

func loadData() {
	_ = controller.LoadFromJSON_Controller()
}

func menuCRUDUser() {
	for {
		fmt.Println()
		fmt.Println("===== MANAJEMEN DATA USER =====")
		fmt.Println("1. Insert User")
		fmt.Println("2. View User")
		fmt.Println("3. Kembali")

		var pilih int
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			view.UserInsertView()

		case 2:
			view.UserViewAll()

		case 3:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func menuCRUDPaket() {
	for {
		fmt.Println()
		fmt.Println("===== MANAJEMEN PAKESET =====")
		fmt.Println("1. Insert Paket")
		fmt.Println("2. Update Paket")
		fmt.Println("3. Delete Paket")
		fmt.Println("4. View Paket")
		fmt.Println("5. Kembali")

		var pilih int
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			view.PaketInsertView()

		case 2:
			view.PaketUpdateView()

		case 3:
			view.PaketDeleteView()

		case 4:
			view.PaketViewAll()

		case 5:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
func menuCRUDMenu() {

	for {
		fmt.Println()
		fmt.Println("===== MANAJEMEN MENU =====")
		fmt.Println("1. Insert Menu")
		fmt.Println("2. Update Menu")
		fmt.Println("3. Delete Menu")
		fmt.Println("4. View Menu")
		fmt.Println("5. Kembali")

		var pilih int
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			view.MenuInsertView()

		case 2:
			view.MenuUpdateView()

		case 3:
			view.MenuDeleteView()

		case 4:
			view.MenuViewAll()

		case 5:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func menuCRUDTransaksi() {
	for {
		fmt.Println()
		fmt.Println("===== MANAJEMEN TRANSAKSI =====")
		fmt.Println("1. Insert Transaksi")
		fmt.Println("2. Update Transaksi")
		fmt.Println("3. Delete Transaksi")
		fmt.Println("4. View Transaksi")
		fmt.Println("5. Kembali")

		var pilih int
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			view.TransaksiInsertView()

		case 2:
			fmt.Println("menu belum tersedia")

		case 3:
			fmt.Println("menu belum tersedia")

		case 4:
			view.TransaksiViewAll()

		case 5:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func ProgramUtama() {
	for {
		fmt.Println()
		fmt.Println("===== MENU UTAMA =====")
		fmt.Println("1. Manajemen Data User")
		fmt.Println("2. Manajemen Data Menu")
		fmt.Println("3. Manajemen Paket")
		fmt.Println("4. Transaksi Penjualan")
		fmt.Println("5. Keluar")

		var pilih int
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			menuCRUDUser()

		case 2:
			menuCRUDMenu()

		case 3:
			menuCRUDPaket()

		case 4:
			menuCRUDTransaksi()

		case 5:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
func main() {
	loadData()

	// ProgramUtama()

	// //view html
	http.HandleFunc("/login", handlers.ShowLoginHandler)
	http.HandleFunc("/login-process", handlers.LoginHandler)
	http.HandleFunc("/success", handlers.SuccessHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/users", handlers.UserPageHandler)
	http.HandleFunc("/users/create", handlers.UserCreatePageHandler)
	http.HandleFunc("/users/store", handlers.UserStoreHandler)
	http.HandleFunc("/users/delete", handlers.UserDeleteHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
