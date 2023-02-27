package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/data?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	} else {
		log.Println("Database connection established")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable() {
	db := OpenConnection()

	var createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
	id SERIAL PRIMARY KEY,
	userId INT NOT NULL,
	title TEXT,
	body TEXT
	);
	`
	_, err := db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create table", err)
	}
	defer db.Close()
}

func main() {

	r := chi.NewRouter()
	CreateTable()
	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root route"))
	})

	r.Get("/posts", GetAllPosts)
	r.Post("/posts", AddPost)
	r.Get("/posts/{id}", GetPostByID)

	fmt.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

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
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}

func AddPost(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO posts (userid, title, body) VALUES ($1, $2, $3)", post.UserId, post.Title, post.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	idStr := chi.URLParam(r, "id")
	fmt.Print(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM posts WHERE userid = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	defer rows.Close()

	if len(posts) == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
