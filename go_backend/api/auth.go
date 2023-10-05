package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	db "Auto_Bangumi/v2/database"
	models "Auto_Bangumi/v2/models"
)

func authError(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(models.NewResponseModel(401, "User not found", "用户不存在"))
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(response)
}

func authReuse(r *http.Request) bool {
	for _, cookie := range r.Cookies() {
		if cookie.Name == "loginSession" {
			_, err := db.FindOne("sessions", "username", cookie.Value)
			if err == nil {
				return true
			}
		}
	}

	return false

}

func AuthLogin(w http.ResponseWriter, r *http.Request) {

	// Check if already logged in
	if authReuse(r) {
		response, _ := json.Marshal(models.NewResponseModel(200, "Logged in successfully", "登陆成功"))
		w.Write(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		authError(w, r)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		authError(w, r)
		return
	}

	admin, err := db.FindOne("users", "username", username)
	if err != nil {
		authError(w, r)
		return
	}

	res := bcrypt.CompareHashAndPassword(admin.Get("password").([]byte), []byte(password))
	if res != nil {
		authError(w, r)
		return
	}

	// Store login session for reuse
	session_uuid := uuid.New().String()
	cookie := http.Cookie{}
	cookie.Name = "loginSession"
	cookie.Value = session_uuid
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	cookie.Path = "/api/v1/auth/login"
	http.SetCookie(w, &cookie)

	// Insert session into database
	if _, err := db.InsertOne("sessions", username, session_uuid, 7*24*time.Hour); err != nil {
		log.Fatal().Msg(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// return 200 with response model
	response, _ := json.Marshal(models.NewResponseModel(200, "Logged in successfully", "登陆成功"))

	w.Write(response)
}
