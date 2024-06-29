package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/database"
	"github.com/Shubhangcs/go-water-dispenser/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	db := database.NewConnection().Db
	defer db.Close()

	// Create a mutex for handling concurrent database access
	var mut sync.Mutex

	// Create a new router instance
	router := mux.NewRouter()

	// Logger middleware (if needed)
	// router.Use(middlewares.LoggerMiddleware)

	// Setup routes
	routes.TransactionRouter(db, &mut, router)

	// Serve static files if needed
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
