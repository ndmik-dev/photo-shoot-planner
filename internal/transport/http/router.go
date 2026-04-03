package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ndmik-dev/photo-shoot-planner/internal/shoot"
)

func NewRouter(shootHandler *shoot.Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	r.Route("/api/v1/shoots", func(r chi.Router) {
		r.Post("/", shootHandler.Create)
		r.Get("/", shootHandler.List)
		r.Get("/{id}", shootHandler.GetByID)
		r.Put("/{id}", shootHandler.Update)
		r.Patch("/{id}/status", shootHandler.PatchStatus)
		r.Delete("{id}", shootHandler.Delete)
	})

	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
