package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DbInstance struct {
	Db *sql.DB
}

func NewConnection() *DbInstance {
	db, err := sql.Open("postgres", "user=root dbname=water sslmode=disable password=password host=localhost")
	if err != nil {
		panic(err)
	}
	return &DbInstance{
		Db: db,
	}
}
