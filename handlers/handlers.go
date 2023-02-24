package handlers

import (
	"myapi/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var DBInstance db.Database

func NewHandler(db db.Database) http.Handler {
	r := chi.NewRouter()
	DBInstance = db
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	return r
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	render.Render(w, r, ErrNotFound)
}
