package handler

import (
	"go-tailwind/view/healthcheck"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, healthcheck.Index())
}
