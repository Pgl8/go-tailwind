package handler

import (
	"go-tailwind/view/healthcheck"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) error {
	healthcheck.Index().Render(r.Context(), w)
	return nil
}
