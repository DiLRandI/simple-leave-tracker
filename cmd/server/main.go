package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"simple-leave-tracker/internal/app"
	"simple-leave-tracker/internal/app/modules/health"
	"simple-leave-tracker/internal/storage/db"
)

var version string

func main() {
	_, err := db.New(DbDSN())
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).
		With("version", version)

	logger.InfoContext(ctx, "starting the server")
	defer logger.InfoContext(ctx, "stopping the server")

	mux := http.NewServeMux()
	app := app.New(&app.Config{
		Log: logger,
		Modules: []app.RouteRegister{
			health.New(logger),
		},
	})
	app.RegisterRoutes(mux)

	server := &http.Server{
		Handler:  mux,
		Addr:     ":8080",
		ErrorLog: log.New(os.Stdout, "[SERVER]", log.LstdFlags),
	}
	defer server.Shutdown(ctx)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	<-ctx.Done()
}

func DbDSN() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)

	return dsn
}
