package db

import "fmt"

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")
