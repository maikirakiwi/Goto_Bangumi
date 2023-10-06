package main

import (
	"fmt"
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
	port, exists := db.Cache.Get("config")
	if exists {
		// Convert cryptic interface{} to string
		port = fmt.Sprintf("%.0f", port.(map[string]interface{})["program"].(map[string]interface{})["webui_port"].(float64))
	} else {
		port = "7892"
	}
	log.Info().Msgf("GotoBangumi Listening on %s", host_ip()+":"+port.(string))
	log.Fatal().Msg(http.ListenAndServe(host_ip()+":"+port.(string), api.Router()).Error())

	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")

}
