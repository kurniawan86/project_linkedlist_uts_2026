package controllers

import (
	"errors"
	"project_linkedlist_uts_2026/domain"
	"project_linkedlist_uts_2026/models"
	"time"
)

func UserInsertController(nama string, dateUpdated string, role string) (bool, error) {
	lastId := models.GetLastID()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	timeCreated := now.Format("2006-01-02 15:04")

	roles := []string{"admin", "kasir", "pelanggan"}

	if nama == "" || role == "" {
		return false, errors.New("nama dan role tidak boleh kosong")
	}

	cek := false
	for _, r := range roles {
		if r == role {
			cek = true
		}
	}
	if cek == false {
		return false, errors.New("role tidak valid")
	}

	datauser := models.InsertModelUser(lastId+1, timeCreated, "", nama, role)
	if datauser.DataUser.IdUser != 0 {
		models.SaveToJSON()
		return true, nil
	}

	err := models.SaveToJSON()
	if err == nil {
		return true, nil
	}
	return false, err
}

func ShowAllUserController() []domain.UserNode {
	dataUsers := models.ShowAllModelUser()
	if dataUsers != nil {
		return dataUsers
	}
	return nil
}

func LoadFromJSON_Controller() error {
	return models.LoadFromJSON()
}

func DeleteUserController(id int) (bool, error) {
	models.DeleteUserModelById(id)
	models.SaveToJSON()
	return true, nil
}

func UpdateUserController(id int, nama string, role string) (bool, error) {
	models.UpdateUserModelById(id, nama, role)
	models.SaveToJSON()
	return true, nil
}

func LoginController(nama string, password string) (domain.UserNode, error) {
	user := models.GetUserByName(nama)
	if user.IdUser == 0 {
		return user, errors.New("user tidak ditemukan")
	}
	if "password" != password {
		return user, errors.New("password salah")
	}
	return user, nil
}
