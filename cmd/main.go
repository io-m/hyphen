package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	di "github.com/io-m/hyphen/internal/dependency"

	"github.com/io-m/hyphen/internal/shared/config"
	db_connection "github.com/io-m/hyphen/internal/shared/db"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

func main() {
	helpers.LoadEnv(constants.DEV_CONFIG_FILE)
	helpers.InitStructuredLogger()

	postgresConnection, err := db_connection.NewPostgresConnection()
	if err != nil {
		slog.Error(err.Error())
	}
	// Deferring Postgres conn closing
	defer func() {
		if err := postgresConnection.Close(); err != nil {
			slog.Error(err.Error())
		}
	}()

	redisClient, err := db_connection.CreateRedisConnection()
	if err != nil {
		slog.Error(err.Error())
	}
	// Deferring Redis conn closing
	defer func() {
		if err := redisClient.Close(); err != nil {
			slog.Error(err.Error())
		}
	}()

	// Global singleton
	appConfig := config.NewAppConfig(postgresConnection, redisClient)

	port := os.Getenv(constants.APP_PORT)
	slog.Info("listening on port: %s............\n", port)
	di.ConfigureRoutes(appConfig)

	// Run server in separate goroutine
	go func() {
		runServer(port, appConfig.GetMux())
	}()

	// Handling graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	slog.Info("Shutting down gracefully...")
}

// Function for running server
func runServer(port string, mux *chi.Mux) {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		slog.Error("Server is down!", err.Error())
	}
}
