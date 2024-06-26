package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/database"
	"github.com/Shubhangcs/go-water-dispenser/middlewares"
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
	-> sync
*/

func main() {
	//Instance to be passed as dependency
	db := database.NewConnection()
	router := mux.NewRouter()
	mut := sync.Mutex{}

	//Middlewares to be used
	router.Use(middlewares.LoggerMiddleware)

	//Controllers to be used
	routes.TransactionRouter(db.Db , &mut , router)

	//Server Listner Implamentation
	var PORT string = ":8000"
	log.Println("Server is running at PORT:", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
