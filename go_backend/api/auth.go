package auth

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// TODO use db abstraction to check credentials
}
