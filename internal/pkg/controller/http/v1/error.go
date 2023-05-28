package v1

import (
	"net/http"

	"github.com/go-chi/render"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(resp http.ResponseWriter, req *http.Request, code int, msg string) {
	resp.WriteHeader(code)
	render.JSON(resp, req, response{msg})
}
