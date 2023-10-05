package main

import (
	"net/http"
	"os"

	"github.com/flosch/pongo2/v6"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"Auto_Bangumi/v2/api"
	db "Auto_Bangumi/v2/database"
)

func host_ip() string {
	_, v6 := os.LookupEnv("IPV6")
	ip, ip_set := os.LookupEnv("HOST")

	if v6 {
		return "::"
	} else if ip_set {
		return ip
	} else {
		return "0.0.0.0"
	}

}

// Faster to compile first at startup
var index = pongo2.Must(pongo2.FromFile("./dist/index.html"))

func templater(w http.ResponseWriter, r *http.Request) {
	err := index.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Database Setup
	go db.Init()
	defer db.Conn.Close()

	// Routing Setup
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	router := chi.NewRouter()

	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist/assets"))))
	router.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./dist/images"))))
	router.Post("/api/v1/auth/login", api.AuthLogin)
	router.Get("/", templater)
	log.Fatal().Msg(http.ListenAndServe(host_ip()+":7892", router).Error())
	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")

}
