package models

import (
	"encoding/json"
	"errors"
	"os"
	"project_linkedlist_uts_2026/domain"
)

var databaseUser domain.LL_User

func InsertModelUser(id int, dataCreated string, dataUpdate string, nama string, role string) domain.LL_User {
	dataUser := domain.UserNode{
		IdUser:      id,
		DateCreated: dataCreated,
		DateUpdated: dataUpdate,
		NamaUser:    nama,
		Role:        role,
	}

	newUserNode := domain.LL_User{
		DataUser: dataUser,
		Linked:   nil,
	}

	if databaseUser.Linked == nil {
		databaseUser.Linked = &newUserNode
		return newUserNode
	}

	var tempLL *domain.LL_User
	tempLL = &databaseUser
	for tempLL.Linked != nil {
		tempLL = tempLL.Linked
	}
	tempLL.Linked = &newUserNode
	return newUserNode
}

func ShowAllModelUser() []domain.UserNode {
	var tabelUser []domain.UserNode
	var tempLL *domain.LL_User
	tempLL = &databaseUser

	for tempLL.Linked != nil {
		tabelUser = append(tabelUser, tempLL.Linked.DataUser)
		tempLL = tempLL.Linked
	}
	return tabelUser
}

func UpdateUserModelById(id int, nama string, role string) bool {
	tempLL := &databaseUser

	for tempLL.Linked != nil {
		if id == tempLL.Linked.DataUser.IdUser {
			tempLL.Linked.DataUser.NamaUser = nama
			tempLL.Linked.DataUser.Role = role
			return true
		}
		tempLL = tempLL.Linked
	}
	return false
}

func DeleteUserModelById(id int) bool {
	tempLL := &databaseUser

	for tempLL.Linked != nil {
		if id == tempLL.Linked.DataUser.IdUser {
			tempLL.Linked = tempLL.Linked.Linked
			return true
		}
		tempLL = tempLL.Linked
	}
	return false
}

func GetLastID() int {
	var tabelUser []domain.UserNode
	var tempLL *domain.LL_User
	tempLL = &databaseUser

	for tempLL.Linked != nil {
		tabelUser = append(tabelUser, tempLL.Linked.DataUser)
		tempLL = tempLL.Linked
	}
	return tempLL.DataUser.IdUser
}

func SaveToJSON() error {
	tableUser := ShowAllModelUser()

	file, err := os.Create("database/user.json")
	if err != nil {
		return errors.New("data tidak berhasil disimpan pada storage")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tableUser)
	return errors.New("data berhasil tersimpan di storage")
}

func LoadFromJSON() error {
	file, err := os.Open("database/user.json")
	if err != nil {
		return errors.New("data gagal dimuat")
	}
	defer file.Close()
	var data []domain.UserNode
	json.NewDecoder(file).Decode(&data)

	databaseUser.Linked = nil

	for _, v := range data {
		InsertModelUser(
			v.IdUser,
			v.DateCreated,
			v.DateUpdated,
			v.NamaUser,
			v.Role,
		)
	}
	return nil
}

func GetUserById(id int) domain.UserNode {
	tempLL := &databaseUser

	for tempLL.Linked != nil {
		if id == tempLL.Linked.DataUser.IdUser {
			return tempLL.Linked.DataUser
		}
		tempLL = tempLL.Linked
	}
	return domain.UserNode{}
}

func GetUserByName(nama string) domain.UserNode {
	tempLL := &databaseUser

	for tempLL.Linked != nil {
		if nama == tempLL.Linked.DataUser.NamaUser {
			return tempLL.Linked.DataUser
		}
		tempLL = tempLL.Linked
	}
	return domain.UserNode{}
}
