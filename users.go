package main

type Users struct {
	userId int    `json:"userid"`
	id     int    `json:"id"`
	title  string `json:"title"`
	body   string `json:"body"`
}

type Getter interface {
	GetAll() []Users
}
