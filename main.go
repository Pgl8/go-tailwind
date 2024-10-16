package main

import (
	"go-tailwind/handler"
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

	router.Handle("/*", public())
	router.Get("/healthcheck", handler.Make(handler.HandleHealthCheck))
	listenPort := os.Getenv("LISTEN_PORT")
	slog.Info("Server is running", "listenAddr", listenPort)
	http.ListenAndServe(listenPort, router)
}