package handlers

import (
	"myapi/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	r := chi.NewRouter()
	dbInstance = db
	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)
	return r
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
