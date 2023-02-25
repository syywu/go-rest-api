package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var db *sql.DB

func main() {

	var err error
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/data?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	var createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
	id SERIAL PRIMARY KEY,
	userId VARCHAR(100) NOT NULL,
	title TEXT,
	body TEXT
	);
	`

	_, err = db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create table", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root route"))
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", GetAllPosts)
		r.Post("/", AddPost)
	})

	fmt.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetAllPosts() {

}

func AddPost() {

}
