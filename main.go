package main

import (
	"/Users/samantha/projects/go-rest-api/posts/posts.go"
	"fmt"
	"log"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// post
func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("posted"))
}

// delete

// patch

// put

func main() {
	new := posts.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// get
	r.Get("/", homeHandler)
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		p := new.GetAll()

	})

	r.Post("/users", postHandler)
	fmt.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
