package models

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

func (l *List) GetPosts() []Post {
	return l.Posts
}

func (l *List) Add(post Post) {
	l.Posts = append(l.Posts, post)
}
