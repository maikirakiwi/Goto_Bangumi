package api

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"
	"golang.org/x/crypto/bcrypt"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
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

func setTokenCookie(w http.ResponseWriter, r *http.Request, token string) {
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = token
	cookie.MaxAge = 7 * 24 * 60 * 60
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	activeUser, exists := db.Cache.Get("activeUser")
	if !exists {
		writeException(w, r, 500, "Internal Server Error")
		return
	}
	_, token, _ := createAccessToken(activeUser.(string))
	setTokenCookie(w, r, token)

	writeJWTResponse(w, r, token, "Bearer", "")
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

	userdata := models.User{}
	dbres := db.Conn.Where(&models.User{Username: username}).First(&userdata)
	if dbres.Error != nil {
		randomSleep()
		writeResponse(w, r, 401, "User not found", "用户不存在")
		return
	}

	res := bcrypt.CompareHashAndPassword(userdata.Password, []byte(password))
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
	cookie.Path = "/"
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	db.Cache.Delete("activeUser")

	//writeResponse(w, r, 200, "Logged out successfully", "登出成功")

	// Blame frontend for not handling this properly
	writeJWTResponse(w, r, "", "", "logout success")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var form map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil || len(form["username"].(string)) > 20 || len(form["username"].(string)) < 4 || len(form["password"].(string)) < 8 {
		writeException(w, r, 500, "Internal Server Error")
		return
	}
	//old_user, exists := db.Cache.Get("activeUser")
	//if !exists {
	//writeException(w, r, 500, "Internal Server Error")
	//return
	//}
	if activeUser, exists := db.Cache.Get("activeUser"); exists && activeUser != form["username"].(string) {
		db.Cache.SetDefault("activeUser", form["username"].(string))
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(form["password"].(string)), bcrypt.DefaultCost)

	// DB transaction
	user := models.User{}
	db_res := db.Conn.First(&user)
	if db_res.Error != nil {
		log.Error().Msg(db_res.Error.Error())
		writeException(w, r, 500, "Internal Server Error")
		return
	}
	user.Username = form["username"].(string)
	user.Password = password
	db.Conn.Save(&user)

	// Set cache and reissue jwt token
	new_user, _ := db.Cache.Get("activeUser")
	_, token, _ := createAccessToken(new_user.(string))
	setTokenCookie(w, r, token)
	writeJWTResponse(w, r, token, "Bearer", "update success")
}
