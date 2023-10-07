package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
	start := time.Now()

	// Database Setup
	db.Init()
	defer db.Teardown()
	log.Info().Msg("Database Initialized.")

	// Routing Setup
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	port, exists := db.Cache.Get("config")
	if exists {
		// Convert cryptic interface{} to string
		port = fmt.Sprintf("%d", port.(map[string]interface{})["program"].(map[string]interface{})["webui_port"].(int64))
	} else {
		port = "7892"
	}

	elapsed := time.Since(start)
	log.Info().Msgf(`GotoBangumi Initialized in %s. Listening on %s`, elapsed.String(), host_ip()+":"+port.(string)+".")
	log.Fatal().Msg(http.ListenAndServe(host_ip()+":"+port.(string), api.Router()).Error())

	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")

}
