package db

import (
	"myapi/models"
)

func (db Db) GetAllPosts() (*models.List, error) {
	list := &models.List{}
	rows, err := db.Conn.Query("SELECT * FROM posts")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			return list, err
		}
		list.Posts = append(list.Posts, post)
	}
	return list, nil
}

func (db Db) AddItem(post *models.Post) error {
	var id int
	query := `INSERT INTO posts (userId, title, body) VALUES ($!, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, post.UserId, post.Title, post.Body).Scan(&id)
	if err != nil {
		return err
	}
	post.Id = id
	return nil
}
