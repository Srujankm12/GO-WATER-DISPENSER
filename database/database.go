package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type databaseConnection struct {
	Connection *sql.DB
}

func NewDatabaseConnectionInstance() *databaseConnection {
	db, err := sql.Open("postgres", "user=root dbname=water sslmode=disable password=password host=localhost")
	if err != nil {
		panic(err)
	}
	return &databaseConnection{
		Connection: db,
	}
}
