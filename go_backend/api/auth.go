package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "Auto_Bangumi/v2/database"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	if user, err := admin_entry; err != nil {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	} else {
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
			http.Error(w, "Invalid login", http.StatusUnauthorized)
			return
		}
	}
}
