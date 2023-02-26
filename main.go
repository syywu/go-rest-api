package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

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
	db, err := sql.Open("postgres", "postgres://user:password@localhost/data?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	} else {
		log.Println("Database connection established")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
	id SERIAL PRIMARY KEY,
	userId INT NOT NULL,
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

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []Post{}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO posts (userid, title, body) VALUES ($1, $2, $3) RETURNING id", post.UserId, post.Title, post.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
