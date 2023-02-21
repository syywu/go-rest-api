package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myapi/posts"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// posts

// {
//     "userId": 1,
//     "id": 1,
//     "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
//     "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
//   },

// delete

// patch

// put

func main() {
	new := posts.New()
	new.Add(posts.Post{
		UserId: "100",
		Id:     "1",
		Title:  "hello",
		Body:   "hello world",
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Route"))
	})

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		p := new.GetPosts()
		json.NewEncoder(w).Encode(p)

	})

	// post
	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posted"))
	})
	fmt.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
