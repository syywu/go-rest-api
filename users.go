package main

type User struct {
	userId int    `json:"userid"`
	id     int    `json:"id"`
	title  string `json:"title"`
	body   string `json:"body"`
}

type Getter interface {
	GetAll() []User
}

type Adder interface {
	Add(user User)
}

type List struct {
	Users []User
}

func New() *List {
	return &List{
		Users: []User{},
	}
}
