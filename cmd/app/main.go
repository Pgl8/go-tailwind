package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}
	router := chi.NewMux()

	router.Get("/healthcheck", handleHealthcheck())
	listenPort := os.Getenv("LISTEN_PORT")
	slog.Info("Server is running", "listenAddr", listenPort)
	http.ListenAndServe(listenPort, router)
}

func handleHealthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Everything is fine!"))
	}
}
