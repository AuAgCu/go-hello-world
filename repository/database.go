package repository

import (
	"database/sql"
	"os"
)

var Db *sql.DB

func init() {
	var err error

	Db, err = sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
