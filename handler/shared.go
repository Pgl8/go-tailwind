package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPhandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Error handling request", "error", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
