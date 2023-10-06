package api

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	db "Auto_Bangumi/v2/database"
	models "Auto_Bangumi/v2/models"
)

var jwtKey = generateKey()
var jwtInstance = jwtauth.New("HS256", jwtKey, jwtKey)

func generateKey() []byte {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal().Msg(err.Error())
	}
	return bytes
}

func writeResponse(w http.ResponseWriter, r *http.Request, statusCode int, msg_en string, msg_zh string) {
	response, _ := json.Marshal(models.NewResponseModel(statusCode, msg_en, msg_zh))
	w.WriteHeader(statusCode)
	w.Write(response)
}

func writeJWTResponse(w http.ResponseWriter, r *http.Request, token string, token_type string, message string) {
	response, _ := json.Marshal(models.NewJWTModel(token, token_type, message))
	w.Write(response)
}

func writeException(w http.ResponseWriter, r *http.Request, statusCode int, detail string) {
	response, _ := json.Marshal(models.NewExceptionModel(statusCode, detail))
	w.WriteHeader(statusCode)
	w.Write(response)
}

// Customized from https://github.com/go-chi/jwtauth/blob/v5.1.1/jwtauth.go#L161
func verifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if it's even a valid jwt
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			writeException(w, r, 401, "Unauthorized")
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			writeException(w, r, 401, "Unauthorized")
			return
		}

		// Check if the user is in session
		activeUser, exists := db.Cache.Get("activeUser")
		if !exists {
			writeException(w, r, 500, "Internal Server Error")
			return
		}

		if username, exists := token.Get("sub"); !exists || username == "" || activeUser.(string) != username.(string) {
			writeException(w, r, 401, "Unauthorized")
			return
		}

		// Token is now authenticated
		next.ServeHTTP(w, r)
	})
}

func createAccessToken(username string) (jwt.Token, string, error) {
	return jwtInstance.Encode(map[string]interface{}{"sub": username, "exp": time.Now().Add(7 * 24 * time.Hour).Unix()})
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	activeUser, exists := db.Cache.Get("activeUser")
	if !exists {
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	cookie := http.Cookie{}
	cookie.Name = "token"
	_, cookie.Value, _ = createAccessToken(activeUser.(string))
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	writeJWTResponse(w, r, cookie.Value, "Bearer", "")
}

func randomSleep() {
	duration, _ := rand.Int(rand.Reader, big.NewInt(1000))
	time.Sleep(time.Duration(duration.Int64()) * time.Millisecond)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// Check if already logged in
	jwt, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		if username, exists := jwt.Get("sub"); exists {
			if activeUser, exists := db.Cache.Get("activeUser"); exists && activeUser == username.(string) {
				writeResponse(w, r, 200, "Logged in successfully", "登陆成功")
				return
			}
		}
	}

	err = r.ParseForm()
	if err != nil {
		randomSleep()
		writeResponse(w, r, 401, "User not found", "用户不存在")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		randomSleep()
		writeResponse(w, r, 401, "User not found", "用户不存在")
		return
	}

	admin, err := db.FindOne("users", "username", username)
	if err != nil {
		randomSleep()
		writeResponse(w, r, 401, "User not found", "用户不存在")
		return
	}

	res := bcrypt.CompareHashAndPassword(admin.Get("password").([]byte), []byte(password))
	if res != nil {
		randomSleep()
		writeResponse(w, r, 401, "User not found", "用户不存在")
		return
	}

	// Store login session for reuse
	db.Cache.SetDefault("activeUser", username)
	refreshTokenHandler(w, r)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	db.Cache.Delete("activeUser")

	//writeResponse(w, r, 200, "Logged out successfully", "登出成功")

	// Blame frontend for not handling this properly
	writeJWTResponse(w, r, "", "", "logout success")
}
