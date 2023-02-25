package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialise(username, password, database string) (Database, error) {
	db := Database{}
	// dsn := fmt.Sprintf("host=%v port=%v user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", "postgres://user:password@localhost:5432/data?sslmode=disable")
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
