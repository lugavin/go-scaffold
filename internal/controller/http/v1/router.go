// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// Swagger docs.
	"github.com/lugavin/go-easy/internal/usecase"
	"github.com/lugavin/go-easy/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(router *chi.Mux, l logger.Interface, t usecase.Translation) {
	// Options
	router.Use(middleware.Logger)

	// Swagger
	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//router.Get("/swagger/*any", swaggerHandler)

	// K8s probe
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Prometheus metrics
	//handler.Get("/metrics", promhttp.Handler())

	// Routers
	router.Route("/v1", func(r chi.Router) {
		newTranslationRoutes(r, t, l)
	})
}
