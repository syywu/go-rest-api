package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text`
	Message    string `json:"message`
}

var (
	ErrMethodNotAllowed = &ErrResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrResponse{StatusCode: 404, Message: "Not found"}
	ErrBadRequest       = &ErrResponse{StatusCode: 400, Message: "Bad request"}
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
