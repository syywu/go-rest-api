package db

import (
	"database/sql"
	"log"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/data?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	} else {
		log.Println("Database connection established")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable() {
	db := OpenConnection()

	var createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
	id SERIAL PRIMARY KEY,
	userId INT NOT NULL,
	title TEXT,
	body TEXT
	);
	`
	_, err := db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create table", err)
	}
	defer db.Close()
}
