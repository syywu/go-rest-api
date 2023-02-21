package posts

type Post struct {
	UserId string `json:"userid"`
	Id     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
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

func (l *List) GetPosts() []Post {
	return l.Posts
}

func (l *List) Add(post Post) {
	l.Posts = append(l.Posts, post)
}
