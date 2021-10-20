package repository

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error

	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
