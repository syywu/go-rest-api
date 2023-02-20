package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// posts
/*
{
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  },*/

type Users struct {
	userId int    `json:"userid"`
	id     int    `json:"id"`
	title  string `json:"title"`
	body   string `json:"body"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/users", usersHandler)
	fmt.Print("listening on port 8080")
	http.ListenAndServe(":8080", r)
}
