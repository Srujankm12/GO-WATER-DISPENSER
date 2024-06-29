package routes

import (
	"database/sql"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/controllers"
	"github.com/Shubhangcs/go-water-dispenser/repository"
	"github.com/gorilla/mux"
)

func TransactionRouter(sql *sql.DB, mut *sync.Mutex, router *mux.Router) {
	repo := repository.NewTransactionRepositoryInstance(sql, mut)
	controller := controllers.NewTransactionControllerInstance(repo)

	// Transaction confirmation route
	router.HandleFunc("/transaction", controller.ConfirmTransactionImpl).Methods("POST")
	//quantity
	router.HandleFunc("/quantity", controller.GetQuantityImpl).Methods("GET")
}
