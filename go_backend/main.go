package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"Auto_Bangumi/v2/api"
	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
)

func host_ip() string {
	_, v6 := os.LookupEnv("IPV6")
	ip, ip_set := os.LookupEnv("HOST")

	if v6 {
		return "::"
	}
	if ip_set {
		return ip
	}
	return "0.0.0.0"

}

var server *http.Server

// Graceful shutdown handler
func main() {
	logFile, _ := os.OpenFile(
		"log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stderr, logFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go Init()

	sig := <-signalCh
	log.Warn().Msgf("GotoBangumi received signal %v. Hammer timeout is set to 30 seconds.", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Catch other errors
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("GotoBangumi failed to gracefully shutdown: %v", err)
	}

	log.Info().Msg("GotoBangumi has gracefully shutdown.")
	api.ClearLog()
}

func Init() {
	start := time.Now()

	// Set logging pref
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Database Setup
	db.Init()
	defer db.Teardown()

	log.Info().Msg("[1/1] Database Initialized.")

	// Qbittorrent Setup
	//downloaders.Init()
	cfg, exists := db.Cache.Get("config")

	// Default port
	port := "7892"
	if exists {
		port = fmt.Sprintf("%d", cfg.(models.Config).Program.WebuiPort)
	}

	// Register server parameters
	server = &http.Server{
		Addr:    host_ip() + ":" + port,
		Handler: api.Router(),
	}

	log.Info().Msgf("GotoBangumi initialized in %s. Listening on %s. ", time.Since(start).String(), server.Addr)
	log.Info().Msgf("GitHub: github.com/maikirakiwi/Goto_Bangumi")
	log.Info().Msgf("Authors: EstrellaXD Twitter: @Estrella_Pan")
	log.Info().Msgf("Authors: Rewrite0 GitHub: @Rewrite0")
	log.Info().Msgf("Authors: Maikiwi Twitter: @notmaikiwi")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Error while starting GotoBangumi: %s", err.Error())
	}
}
