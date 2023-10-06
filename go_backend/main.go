package main

import (
	"net/http"
	"os"

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

func main() {
	// Database Setup
	db.Init()
	defer db.Teardown()
	log.Info().Msg("Database initialized")

	// Routing Setup
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Fatal().Msg(http.ListenAndServe(host_ip()+":7892", api.Router()).Error())

	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")

}
