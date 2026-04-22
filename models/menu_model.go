package models

import (
	"project_linkedlist_uts_2026/domain"
)

var tabelMenu [100]domain.Menu
var indexMenu int = 0

func MenuInsert(nama string, harga int, ket string) bool {
	if nama != "" && harga != 0 {
		if indexMenu == 0 {
			tabelMenu[indexMenu].IdMenu = indexMenu
		} else {
			tabelMenu[indexMenu].IdMenu = (tabelMenu[indexMenu-1].IdMenu) + 1
		}
		tabelMenu[indexMenu].NamaMenu = nama
		tabelMenu[indexMenu].Harga = harga
		tabelMenu[indexMenu].Keterangan = ket
		indexMenu++
		return true
	}
	return false
}

func MenuView() []domain.Menu {
	return tabelMenu[:indexMenu]
}

func MenuUpdate(id int, nama string, harga int, ket string) domain.Menu {
	for i := 0; i < indexMenu; i++ {
		if tabelMenu[i].IdMenu == id {
			if nama != "" {
				tabelMenu[i].NamaMenu = nama
			}
			if harga != 0 {
				tabelMenu[i].Harga = harga
			}
			if ket != "" {
				tabelMenu[i].Keterangan = ket
			}
			return tabelMenu[i]
		}
	}
	return domain.Menu{}
}

func MenuDelete(id int) bool {
	for i := 0; i < indexMenu; i++ {
		if tabelMenu[i].IdMenu == id {
			tabelMenu[i] = tabelMenu[indexMenu-1]
			indexMenu--
			return true
		}
	}
	return false
}

func MenuInsertStatic() {
	MenuInsert("Nasi Goreng", 15000, "Pedas")
	MenuInsert("Mie Ayam", 12000, "Porsi besar")
	MenuInsert("Bakso", 13000, "Bakso urat")
	MenuInsert("Sate Ayam", 20000, "Bumbu kacang")
	MenuInsert("Soto Ayam", 14000, "Kuah bening")
	MenuInsert("Ayam Geprek", 18000, "Level 3")
	MenuInsert("Nasi Uduk", 10000, "Dengan ayam")
	MenuInsert("Lontong Sayur", 11000, "Sayur labu")
	MenuInsert("Gado-Gado", 12000, "Bumbu kacang")
	MenuInsert("Rendang", 25000, "Daging sapi")
}

func GetMenuById(id int) domain.Menu {
	for i := 0; i < indexMenu; i++ {
		if tabelMenu[i].IdMenu == id {
			return tabelMenu[i]
		}
	}
	return domain.Menu{}
}

func GetNamaMenuById(id int) string {
	for i := 0; i < indexMenu; i++ {
		if tabelMenu[i].IdMenu == id {
			return tabelMenu[i].NamaMenu
		}
	}
	return ""
}
