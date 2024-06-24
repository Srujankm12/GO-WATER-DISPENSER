package main

import (
	"log"
	"net/http"
	"github.com/Shubhangcs/go-water-dispenser/database"
	"github.com/Shubhangcs/go-water-dispenser/routes"
	"github.com/gorilla/mux"
)

// This project Mainly Created for a water Dispenser Machine server implementation
// for more info visit: https://github.com/Shubhangcs/go-water-dispenser

//Packages to be used

/*
	-> fmt
	-> log
	-> net/http
	-> json
	-> strings
	-> gorilla mux
	-> sql
	-> sql drivers
*/

func main() {
	db := database.NewDatabaseConnectionInstance()
	router := mux.NewRouter()

	//Controllers Defination
	routes.UserRouter(db.Connection, router)

	//Server Running Code
	var PORT string = ":8000"
	log.Println("Server is running at PORT:", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
