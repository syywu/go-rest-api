package models

import (
	"fmt"
	"net/http"
)

type Post struct {
	UserId int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type List struct {
	Posts []Post
}

func New() *List {
	return &List{
		Posts: []Post{},
	}
}

func (p *Post) Bind(r *http.Request) error {
	if p.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}

// chi,Renderer- render as JSON onjects to API consumers
func (*List) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Post) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (l *List) GetPosts() []Post {
	return l.Posts
}

func (l *List) Add(post Post) {
	l.Posts = append(l.Posts, post)
}
