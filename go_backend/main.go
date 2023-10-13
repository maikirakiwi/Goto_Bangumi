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
	"Auto_Bangumi/v2/models/store"
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

var server *http.Server

// Graceful shutdown handler
func main() {
	// Pretty print by default because we don't care about 1ms of performance each log.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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
}

func Init() {
	start := time.Now()

	// Database Setup
	db.Init()
	defer db.Teardown()
	log.Info().Msg("Database Initialized.")

	// Qbittorrent Setup
	//downloaders.Init()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	cfg, _ := store.BoxForConfigModel(db.Conn).Get(1)

	// Register server parameters
	server = &http.Server{
		Addr:    host_ip() + ":" + fmt.Sprintf("%d", cfg.Program.WebuiPort),
		Handler: api.Router(),
	}

	log.Info().Msgf("GotoBangumi initialized in %s. Listening on %s.", time.Since(start).String(), server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Error while starting GotoBangumi: %s", err.Error())
	}
}
