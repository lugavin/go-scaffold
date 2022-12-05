package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/lugavin/go-scaffold/internal/entity"
	"github.com/lugavin/go-scaffold/internal/usecase"
	"github.com/lugavin/go-scaffold/pkg/logger"
)

type translationRoutes struct {
	t usecase.Translation
	l logger.Interface
}

func newTranslationRoutes(router chi.Router, t usecase.Translation, l logger.Interface) {
	h := &translationRoutes{t, l}

	router.Route("/translation", func(r chi.Router) {
		r.Get("/history", h.history)
		r.Post("/do-translate", h.doTranslate)
	})
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *translationRoutes) history(resp http.ResponseWriter, req *http.Request) {
	translations, err := r.t.History(req.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(resp, req, http.StatusInternalServerError, "database problems")

		return
	}

	render.JSON(resp, req, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
func (r *translationRoutes) doTranslate(resp http.ResponseWriter, req *http.Request) {
	var request doTranslateRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(resp, req, http.StatusBadRequest, "invalid request body")

		return
	}

	translation, err := r.t.Translate(
		req.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(resp, req, http.StatusInternalServerError, "translation service problems")

		return
	}

	render.JSON(resp, req, translation)
}
