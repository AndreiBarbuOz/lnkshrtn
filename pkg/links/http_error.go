package links

import (
	"github.com/go-chi/render"
	"net/http"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Bad request", http.StatusBadRequest)
}

func httpRespondWithError(slug string, err error, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	resp := ErrorResponse{slug, status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}
