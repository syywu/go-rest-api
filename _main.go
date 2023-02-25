package main

import (
	"fmt"
	"log"
	"myapi/db"
	"myapi/handlers"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	// get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root Route"))
	})
	r.Route("/posts", func(r chi.Router) {
		r.Get("/", handlers.GetAllPosts)
		r.Post("/", handlers.CreatePost)
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
	// port := ":8080"
	// listener, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("an error has occured: %s", err)
	// }

	dbUser, dbPassword, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	database, err := db.Initialise(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// ensures db connection is kept on while application is running
	defer database.Conn.Close()

	// API server is started on a seperate goroutine
	// it keeps running until it receives a SIGINT or SIGTERM signal which then calls the Stop func to clean up and shut down server
	// 	httpHandler := handlers.NewHandler(database)
	// 	server := &http.Server{
	// 		Handler: httpHandler,
	// 	}
	// 	go func() {
	// 		server.Serve(listener)
	// 	}()
	// 	defer Stop(server)
	// 	log.Printf("started server on %s", port)
	// 	ch := make(chan os.Signal, 1)
	// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// 	log.Print(fmt.Sprint(<-ch))
	// 	log.Println("Stopping API server")
	// }

	// func Stop(server *http.Server) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	if err := server.Shutdown(ctx); err != nil {
	// 		log.Printf("could not shut down server correctly: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// }

	fmt.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
