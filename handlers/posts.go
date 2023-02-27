package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"myapi/db"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// db := db.OpenConnection()

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	posts := []models.Post{}

	for rows.Next() {
		var post models.Post
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
	db := db.OpenConnection()
	var post models.Post
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
	db := db.OpenConnection()
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

	posts := []models.Post{}
	for rows.Next() {
		post := models.Post{}
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

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()
	params := chi.URLParam(r, "id")
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	row.RowsAffected()
	defer db.Close()

	w.WriteHeader(http.StatusOK)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()
	params := chi.URLParam(r, "id")
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
	}

	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row, err := db.Exec(`UPDATE posts SET userID = $1, title = $2, body = $3 WHERE id = $4`, post.UserId, post.Title, post.Body, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row.RowsAffected()
	defer db.Close()

	w.WriteHeader(http.StatusCreated)
}
