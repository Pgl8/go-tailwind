package handler

import (
	"log/slog"
	"net/http"
)

type HTTPhandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Error handling request", "error", err, "path", r.URL.Path)
		}
	}
}
