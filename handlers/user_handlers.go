package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"project_linkedlist_uts_2026/controllers"
	"project_linkedlist_uts_2026/domain"
	"strconv"
	"strings"
)

func ShowLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view_html/login.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	namaStr := r.FormValue("namauser")
	password := r.FormValue("password")

	user, err := controllers.LoginController(namaStr, password)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:     "user",
		Value:    user.NamaUser + "|" + user.Role,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view_html/dashboard.html")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	parts := strings.Split(cookie.Value, "|")

	nama := "Guest"
	role := "Unknown"

	if len(parts) >= 2 {
		nama = parts[0]
		role = parts[1]
	}

	data := struct {
		Nama string
		Role string
	}{
		Nama: nama,
		Role: role,
	}

	tmpl, err := template.ParseFiles(
		"view_html/dashboard.html",
		"view_html/partials/sidebar.html",
		"view_html/partials/topbar.html",
	)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "user",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // hapus cookie
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func UserPageHandler(w http.ResponseWriter, r *http.Request) {

	// 🔐 cek login (cookie)
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// parse cookie
	parts := strings.Split(cookie.Value, "|")

	nama := "Guest"
	role := "Unknown"

	if len(parts) >= 2 {
		nama = parts[0]
		role = parts[1]
	}

	// ambil data user dari controller
	users := controllers.ShowAllUserController()

	data := struct {
		Nama  string
		Role  string
		Users []domain.UserNode
	}{
		Nama:  nama,
		Role:  role,
		Users: users,
	}

	fmt.Println(data)

	tmpl, err := template.ParseFiles(
		"view_html/users.html",
		"view_html/partials/sidebar.html",
		"view_html/partials/topbar.html",
	)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func UserCreatePageHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	parts := strings.Split(cookie.Value, "|")

	nama := "Guest"
	role := "Unknown"

	if len(parts) >= 2 {
		nama = parts[0]
		role = parts[1]
	}

	data := struct {
		Nama string
		Role string
	}{
		Nama: nama,
		Role: role,
	}

	tmpl, err := template.ParseFiles(
		"view_html/users_create.html",
		"view_html/partials/sidebar.html",
		"view_html/partials/topbar.html",
	)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func UserStoreHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	nama := r.FormValue("nama")
	role := r.FormValue("role")

	// dateUpdated kosong (sesuai controller kamu)
	success, err := controllers.UserInsertController(nama, "", role)
	if !success {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

	// 🔐 hanya izinkan POST
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	// 🔐 cek login
	_, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// ambil id dari form
	idStr := r.FormValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	// panggil controller
	success, err := controllers.DeleteUserController(id)
	if !success {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect balik ke users
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
