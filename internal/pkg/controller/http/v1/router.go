// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lugavin/go-scaffold/internal/pkg/env"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(e *env.Environment) http.Handler {
	router := chi.NewRouter()

	// Options
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	// K8s probe
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Prometheus metrics
	//http.Handle("/metrics", promhttp.Handler())

	// Routers
	router.Route("/v1", func(r chi.Router) {
		newTranslationRoutes(r, e.TranslationUseCase(), e.Logger())
	})

	return router
}
