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
