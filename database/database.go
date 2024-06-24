package database

import "database/sql"



type databaseConnection struct{
	Connection *sql.DB
}

func NewDatabaseConnectionInstance() *databaseConnection {
	db , err := sql.Open("postgres" , "")
	if err != nil {
		panic(err)
	}
	return &databaseConnection{
		Connection: db,
	}
}