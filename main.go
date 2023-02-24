package main

import (
	"fmt"
	"log"
	"myapi/db"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	// var allPosts = []*models.Post{
	// 	{UserId: 100, Id: 1, Title: "hello", Body: "hello world"},
	// }
	// allPosts := models.New()

	// allPosts.Add(models.Post{
	// 	UserId: 100,
	// 	Id:     1,
	// 	Title:  "hello",
	// 	Body:   "hello world",
	// })

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root Route"))
	})

	/*
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", handlers.GetPosts(allPosts))
			// post
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				req := &models.Post{}
				// allPosts = append(allPosts, req)
				// allPosts.Add(models.Post{
				// 	UserId: req["userid"],
				// 	Id:     req["id"],
				// 	Title:  req["title"],
				// 	Body:   req["body"],
				// })
				// req.Id = rand.Intn(100) + 10
				allPosts = append(allPosts, req)
				json.NewDecoder(r.Body).Decode(req)
				w.Write([]byte("posted"))
				render.Status(r, http.StatusCreated)
			})
		})

		r.Route("/{postID}", func(r chi.Router) {
			// delete
			r.Delete("/", handlers.DeletePost())

			// patch

			// put
		})
	*/
	dbUser, dbPassword, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	database, err := db.Initialise(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	// ensures db connection is kept on while application is running
	defer database.Conn.Close()

	fmt.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
