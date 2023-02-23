package db

import (
	"database/sql"
	"fmt"
)

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Db struct {
	Conn *sql.DB
}
