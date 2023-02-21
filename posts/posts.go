package posts

type Post struct {
	userId int    `json:"userid"`
	id     int    `json:"id"`
	title  string `json:"title"`
	body   string `json:"body"`
}

type Getter interface {
	GetAll() []Post
}

type Adder interface {
	Add(post Post)
}

type List struct {
	Posts []Post
}

func New() *List {
	return &List{
		Posts: []Post{},
	}
}

func (l *List) GetUsers() []Post {
	return l.Posts
}

func (l *List) Add(post Post) {
	l.Posts = append(l.Posts, post)
}
